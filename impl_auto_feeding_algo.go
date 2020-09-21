package main

import (
	modle "eabi_gateway/impl"
	"errors"
	"fmt"
	"math"
)

//输出参数
var (
	D_YGJYGD float32 //油管积液高度
	D_TGJYGD float32 //套管积液高度
	D_JYL    float32 //积液量
	D_JiaYL  float32 //
)

func powf(x, y float32) float32 { return float32(math.Pow(float64(x), float64(y))) }
func expf(x float32) float32    { return float32(math.Exp(float64(x))) }
func fabsf(x float32) float32   { return float32(math.Abs(float64(x))) }
func logf(x float32) float32    { return float32(math.Log(float64(x))) }
func sqrtf(x float32) float32   { return float32(math.Sqrt(float64(x))) }

//用Dranchuk-Purvis-Robinson方法计算气体偏差因子
//求解方法-牛顿迭代法
//适用范围-- 1.05《=Tr《=3.0, 0.2《=Pr《=30.0
//<param name="T">管段平均温度，K</param>
//<param name="P">-管段平均压力，MPa</param>
//<param name="Yg">天然气相对密度，无量纲</param>
//<returns>输出参数天然气偏差因子，无量纲</returns>
func z_DPR(T float32, P float32, Yg float32) float32 {

	var djsvaue float32 = 0.0
	// Dim Ppc As Single, Tpc As Single   'Ppc-拟临界压力，Tpc—拟临界温度
	var Ppc float32 = 0.0
	var Tpc float32 = 0.0
	//Dim Ppr As Single, Tpr As Single   'Ppr-拟对比压力，Tpr-拟对比温度
	var Ppr float32 = 0.0
	var Tpr float32 = 0.0
	var a1 float32 = 0.31506237
	var a2 float32 = -1.0467099
	var a3 float32 = -0.57832729
	var A4 float32 = 0.53530771
	var A5 float32 = -0.61232032
	var A6 float32 = -0.10488813
	var A7 float32 = 0.68157001
	var A8 float32 = 0.68446549
	//干气
	Ppc = 4.6677 + 0.1034*Yg - 0.2586*Yg*Yg
	Tpc = 93.3333 + 180.5556*Yg - 6.9444*Yg*Yg

	Ppr = P / Ppc
	Tpr = T / Tpc

	var Rhopr float32 = 1
	var xn1 float32 = Rhopr
	var xn float32 = 0
	var fxn float32 = 0
	var dfxn float32 = 0
	//中间变量
	var Calculate_temp1 float32 = 0
	var Calculate_temp2 float32 = 0
	var Calculate_temp3 float32 = 0
	var Calculate_temp4 float32 = 0
	var Calculate_temp5 float32 = 0
	var Calculate_temp6 float32 = 0

	for {
		xn = xn1

		Calculate_temp1 = -0.27*Ppr/Tpr + xn
		Calculate_temp2 = a1 + a2/Tpr + a3/powf(Tpr, 3)
		Calculate_temp2 *= powf(xn, 2)
		Calculate_temp3 = (A4 + A5/Tpr) * powf(xn, 3)
		Calculate_temp4 = A5 * A6 * powf(xn, 6) / Tpr
		Calculate_temp5 = A7 * (1 + A8*powf(xn, 2))
		Calculate_temp5 *= powf(xn, 3)
		Calculate_temp6 = -A8 * powf(xn, 2)
		Calculate_temp6 = expf(Calculate_temp6)
		Calculate_temp5 *= Calculate_temp6
		Calculate_temp5 = Calculate_temp5 / powf(Tpr, 3)

		fxn = Calculate_temp1 + Calculate_temp2 + Calculate_temp3 + Calculate_temp4 + Calculate_temp5

		Calculate_temp1 = 1 + 2*(a1+a2/Tpr+a3/powf(Tpr, 3))*xn
		Calculate_temp2 = 3 * (A4 + A5/Tpr) * powf(xn, 2)
		Calculate_temp3 = 6 * A5 * A6
		Calculate_temp3 *= powf(xn, 5) / Tpr
		Calculate_temp4 = A7 / powf(Tpr, 3)
		Calculate_temp5 = 3*powf(xn, 4) - 2*A8*powf(xn, 6)
		Calculate_temp5 = 3*powf(xn, 2) + A8*Calculate_temp5
		Calculate_temp6 = -A8 * powf(xn, 2)
		Calculate_temp6 = expf(Calculate_temp6)
		Calculate_temp5 *= Calculate_temp6
		dfxn = Calculate_temp1 + Calculate_temp2 + Calculate_temp3 + Calculate_temp4*Calculate_temp5

		xn1 = xn - fxn/dfxn
		Calculate_temp1 = 1 + (a1+a2/Tpr+a3/(powf(Tpr, 3)))*xn
		Calculate_temp2 = (A4 + A5/Tpr) * powf(xn, 2)
		Calculate_temp3 = A5 * A6 * powf(xn, 5) / Tpr
		Calculate_temp4 = A7 / powf(Tpr, 3)
		Calculate_temp4 *= (1 + A8*powf(xn, 2))
		Calculate_temp5 = powf(xn, 2)
		Calculate_temp6 = -A8 * powf(xn, 2)
		Calculate_temp6 = expf(Calculate_temp6)
		Calculate_temp5 *= Calculate_temp6
		Z = Calculate_temp1 + Calculate_temp2 + Calculate_temp3 + Calculate_temp4*Calculate_temp5

		djsvaue = fabsf(xn1 - xn)
		if (djsvaue / xn) <= 0.0000001 {
			break
		}
	}

	return Z
}

//天然气密度计算模块
//<param name="T">T-井段平均温度，K</param>
//<param name="P">P-井段平均压力，MPa</param>
//<param name="Yg">Yg-天然气相对密度，无量纲</param>
//<returns>输出参数：zρg-天然气密度，kg/m^3</returns>
func f_Rhog(T float32, P float32, Yg float32) float32 {
	var drtnvalue float32 = 0
	var Calculate_temp1 float32 = 0 //中间变量
	//Dim Z As Single
	Z = z_DPR(T, P, Yg)
	Calculate_temp1 = 3484.4 * Yg * P
	drtnvalue = Calculate_temp1 / (Z * T)
	return drtnvalue
}

//天然气粘度计算模块
//<param name="T">T-地层温度，K</param>
//<param name="P">P-地层压力，MPa</param>
//<param name="Yg">Yg-天然气相对密度，无量纲</param>
//<returns>Mug-天然气粘度，mPa*s</returns>
func f_Mug(T float32, P float32, Yg float32) float32 {
	var drtnvalue float32 = 0
	//Mg-混合气相的相对分子质量
	var Mg float32 = 0
	var X float32 = 0
	var Y float32 = 0
	var Rhog float32 = 0
	var k float32 = 0
	//中间变量
	var Calculate_temp1 float32 = 0
	var Calculate_temp2 float32 = 0

	Mg = 28.97 * Yg
	Calculate_temp1 = 2.6832 * powf(10, (-2))
	Calculate_temp1 *= (470 + Mg) * powf(T, 1.5)
	Calculate_temp2 = 116.1111 + 10.5556*Mg + T
	k = Calculate_temp1 / Calculate_temp2
	X = 0.01 * (350 + 54777.78/T + Mg)
	Y = 0.2 * (12 - X)
	//kg/m^3
	Rhog = f_Rhog(T, P, Yg)

	Calculate_temp1 = powf(10, (-4)) * k
	Calculate_temp2 = powf((Rhog / 1000), Y)

	drtnvalue = Calculate_temp1 * Calculate_temp2

	return drtnvalue
}

// 天然气体积系数计算模块
// <param name="T">T-地层温度，K</param>
// <param name="P">P-地层压力，MPa</param>
// <param name="Yg">Yg-天然气相对密度，无量纲</param>
// <returns>Mug-天然气粘度，mPa*s</returns>
func f_Bg(T float32, P float32, Yg float32) float32 {
	var drtnvalue float32 = 0
	var Calculate_temp1 float32 = 0
	var Calculate_temp2 float32 = 0 //中间变量
	Z = z_DPR(T, P, Yg)

	Calculate_temp1 = 3.458 * powf(10, (-4))
	Calculate_temp2 = Z * T / P
	drtnvalue = Calculate_temp1 * Calculate_temp2
	return drtnvalue
}

// 地层水粘度计算模块
// <param name="T">T-温度，K</param>
// <returns>f_Muw-地层水粘度，mPa*s</returns>
func f_Muw(T float32) float32 {
	var rtnvalue float32 = 0
	var dd1 float32 = 0
	var Calculate_temp1 float32 = 0
	//中间变量
	var Calculate_temp2 float32 = 0
	var Calculate_temp3 float32 = 0
	dd1 = 1.8*(T-273) + 32

	Calculate_temp1 = 1.003 - 1.479*powf(10, (-2))*dd1
	Calculate_temp2 = 1.982 * powf(10, (-5))
	Calculate_temp3 = powf(dd1, 2)
	rtnvalue = expf(Calculate_temp1 + Calculate_temp2*Calculate_temp3)

	return rtnvalue
}

//油气表面张力
//<param name="T">T-温度，K</param>
//<param name="P">P-压力，MPa</param>
//<returns>f_σo-油气表面张力，mN/m</returns>
func f_Sigmaw(T float32, P float32) float32 {
	var rtnvalue float32 = 0
	var q137 float32 = 52.5 - 0.87018*P
	var q23 float32 = 76 * expf(-0.0362575*P)
	rtnvalue = (137.78 - (T - 273)) * 1.8 / 206
	rtnvalue = rtnvalue*(q23-q137) + q137
	return rtnvalue
}

// 两相流H_B模型模块
func Hagedorn_Brown() {
	var djsvaue float32 = 0
	var Tav float32 = 0
	//var Pavfloat32 = 0
	var T [10]float32
	var P [10]float32
	var H [10]float32
	var gradP1 [10]float32

	var qwsc, qm, At, dp0, n, dh, gradT, dp, gradP float32
	var Rhog, Bg, Sigmaw, Muw, Mug, Rhol, Vsl, vsg, GL, Gg float32
	var Gm, vm, AA, BB, Ngv, Nlv, Nl, ND, p1, p2 float32
	var p3, p4, p5, p6, p7, p21 float32
	var p22, p23, p24, p25, p26, p27, p28 float32
	var CNl, fX1, fx2, Phi, HL, Rhom, LambdaL, Rhons, Mum, Re float32
	var e, fm, Wm, Vs, Hg, Tauf, hlfai float32
	var Rhow float32 = 1050

	//Rhow 是什么参数   液相密度：默认为1050
	var qgsc float32 = 0 //Qg * powf(10, 4) / 86400; //'产气量，m^3/s
	var i uint32 = 1
	//中间变量
	var Calculate_temp1 float32 = 0
	var Calculate_temp2 float32 = 0
	var Calculate_temp3 float32 = 0
	var Calculate_temp4 float32 = 0
	var Calculate_temp5 float32 = 0

	qgsc = Qg * powf(10, 4)
	qgsc /= 86400
	if Qw == 0 {
		qwsc = 0.1 / 86400
	} else {
		qwsc = Qw / 86400 // '产水量，m^3/s
	}
	qm = qgsc + qwsc                  //
	At = 3.1415926 * powf(Dti, 2) / 4 //
	g = 9.81
	Yg = 0.56
	dp0 = 0.5 //    '定压力增量，MPa
	n = 10
	dh = Ht1 / n
	T[0] = Twh
	P[0] = Pt
	gradT = (Tr - Twh) / Hr

	for i = 1; i < uint32(n); i++ {
		for {
			dp = dp0
			//'1.管段的平均压力Pav和平均温度Tav
			Tav = T[i-1] + gradT*dh/2 //'小段平均温度
			Pav = P[i-1] + dp/2
			//'2.确定平均压力P和平均温度T下的物性参数：Sigmaw=0,Muw=0,Mug=0,Rhol=0
			Z = z_DPR(Tav, Pav, Yg)
			Rhog = f_Rhog(Tav, Pav, Yg)
			Bg = f_Bg(Tav, Pav, Yg)
			Sigmaw = f_Sigmaw(Tav, Pav) / 1000
			Muw = f_Muw(Tav) / 1000
			Mug = f_Mug(Tav, Pav, Yg) / 1000
			Rhol = 1050 // '液相密度，kg/m^3
			//'3.计算气液表观流速  Vsl=0,  vsg=0, GL=0, Gg=0, Gm=0,
			Vsl = qwsc / At      //
			vsg = qgsc * Bg / At //
			GL = At * Vsl * Rhol //'液相质量流量，kg/s
			Gg = At * vsg * Rhog //'气相质量流量，kg/s
			Gm = Gg + GL         //
			//'4.单位流通面积上的混合物流速
			vm = Vsl + vsg //  ' m/(s)
			//'5、适用性判定
			AA = 1.071 - 0.7277*powf((Vsl+vsg), 2)/Dti //
			if AA < 0.13 {
				AA = 0.13
			}
			BB = vsg / (Vsl + vsg) //
			if (BB - AA) >= 0 {    //  If BB - AA > 0 Or BB - AA = 0 Then
				//6.计算4个无因次量 Ngv=0,Nlv=0,Nl=0,ND=0

				Calculate_temp1 = Rhol / (g * Sigmaw)
				Ngv = vsg * powf(Calculate_temp1, 0.25) //气相速度数
				Calculate_temp1 = Rhol / (g * Sigmaw)
				Nlv = Vsl * powf(Calculate_temp1, 0.25) //液相速度数
				Calculate_temp1 = g / (Rhol * powf(Sigmaw, 3))
				Nl = Muw * powf(Calculate_temp1, 0.25) //液相粘度数
				Calculate_temp1 = Rhol * g / Sigmaw
				ND = Dti * powf(Calculate_temp1, 0.5) //管径数
				//7.计算CNl
				p1 = 0.001405082159
				p2 = 0.01580646101
				p3 = -0.136809497344
				p4 = 0.949624462147
				p5 = -2.342844072177
				p6 = 2.486967310526
				p7 = -0.983447366511
				//CNl = p1 + p2 * pow(Nl, 0.5) + p3 * Nl + p4 * pow(Nl, 1.5) + p5 * pow(Nl, 2) + p6 * pow(Nl, 2.5) + p7 * pow(Nl, 3);
				Calculate_temp1 = p1 + p2*powf(Nl, 0.5)
				Calculate_temp2 = p3*Nl + p4*powf(Nl, 1.5)
				Calculate_temp3 = p5 * powf(Nl, 2)
				Calculate_temp4 = p6 * powf(Nl, 2.5)
				Calculate_temp5 = p7 * powf(Nl, 3)
				CNl = Calculate_temp1 + Calculate_temp2 + Calculate_temp3 + Calculate_temp4 + Calculate_temp5

				//'8.计算HL/Phi
				//p11 = 0.093168284
				//p12 = -318.6733901
				//p13 = 100739.677
				//p14 = 32.15203351
				//p15 = -0.000128007
				//fX1 = Nlv * CNl / (pow(Ngv, 0.575) * ND) * pow((Pav / 0.101), 0.1);
				Calculate_temp1 = Nlv * CNl
				Calculate_temp2 = powf(Ngv, 0.575) * ND
				Calculate_temp3 = powf((Pav / 0.101), 0.1)
				fX1 = Calculate_temp1 / Calculate_temp2
				fX1 *= Calculate_temp3
				//'HLFAI = p11 + p12 * fx1 + p13 * fx1 ^ 2.5 + p14 * fx1 ^ 0.5 + p15 / fx1 ^ 0.5
				//hlfai = ((-0.08483253809) + 111.10878774647 * pow(fX1, 0.5) + (-4343.1994122927) * fX1 +
				//583614.509094846 * pow(fX1, 1.5)) / (1 + 271.57873338379 * pow(fX1, 0.5) + (-5907.13288902848) * fX1 +
				//582425.06687108 * pow(fX1, 1.5));
				Calculate_temp1 = (-0.08483253809) + 111.10878774647*powf(fX1, 0.5)
				Calculate_temp1 += (-4343.1994122927) * fX1
				Calculate_temp1 += 583614.509094846 * powf(fX1, 1.5)
				Calculate_temp2 = 1 + 271.57873338379*powf(fX1, 0.5)
				Calculate_temp2 += (-5907.13288902848) * fX1
				Calculate_temp2 += 582425.06687108 * powf(fX1, 1.5)
				hlfai = Calculate_temp1 / Calculate_temp2

				if hlfai < 0 {
					hlfai = 0
				}
				//9.计算Phi
				p21 = 0.92124659726
				p22 = -66.28874853238
				p23 = -51.8665355048
				p24 = 1393.57605439823
				p25 = 524.4716641716
				p26 = -3112.78961898627
				p27 = 17493.270402786
				p28 = 49062.194646086
				fx2 = Ngv * powf(Nl, 0.38)
				fx2 /= powf(ND, 2.14)
				if fx2 < 0.013 {
					Phi = 1
				} else {
					//Phi = (p21 + p23 * fx2 + p25 * pow(fx2, 2) + p27 * pow(fx2, 3)) / (1 + p22 * fx2 + p24 * pow(fx2, 2) + p26 * pow(fx2, 3) + p28 * pow(fx2, 4));
					Calculate_temp1 = p21 + p23*fx2
					Calculate_temp1 += p25 * powf(fx2, 2)
					Calculate_temp1 += p27 * powf(fx2, 3)
					Calculate_temp2 = 1 + p22*fx2 + p24*powf(fx2, 2)
					Calculate_temp2 += p26 * powf(fx2, 3)
					Calculate_temp2 += p28 * powf(fx2, 4)
					Phi = Calculate_temp1 / Calculate_temp2
				}
				//'10.计算HL  HL=0,Rhom=0,LambdaL=0,Rhons=0,Mum=0,Re=0
				HL = hlfai * Phi
				//'11.计算混合物密度
				Rhom = Rhol*HL + Rhog*(1-HL) //混合物密度，kg/m^3
				//'12.计算 Rem
				vm = vsg + Vsl
				LambdaL = Vsl / vm
				Rhons = Rhol*LambdaL + Rhog*(1-LambdaL)
				Mum = powf(Muw, HL)
				Mum *= powf(Mug, (1 - HL))
				Re = Rhons * vm * Dti / Mum
				//'13.用Jain公式计算fm
				e = 0.01524 // '绝对粗糙度，mm e fm
				if Re > 2300 {
					//fm = pow((1.14 - 2 * log(e / (Dti * pow(10, 3)) + 21.25 / pow(Re, 0.9)) / log(10)), (-2));
					Calculate_temp1 = e / (Dti * powf(10, 3))
					Calculate_temp2 = 21.25 / powf(Re, 0.9)
					Calculate_temp3 = 2 * logf(Calculate_temp1+Calculate_temp2)
					Calculate_temp4 = 1.14 - Calculate_temp3/logf(10)
					fm = powf(Calculate_temp4, (-2))
				} else {
					fm = 64 / Re
				}
				//'14.计算压力梯度
				//gradP = Rhom * g + fm * powf(Gm, 2) / (2 * Dti * powf(At, 2) * Rhom);
				Calculate_temp1 = Rhom * g
				Calculate_temp2 = fm * powf(Gm, 2)
				Calculate_temp3 = 2 * Dti * powf(At, 2) * Rhom
				gradP = Calculate_temp1 + Calculate_temp2/Calculate_temp3
				dp0 = gradP * dh * powf(10, -6)
			} else {
				//Wm=0,Vs=0,Hg=0,Tauf=0
				Wm = qgsc*Rhog + qwsc*Rhow
				Vs = 0.244 //  'm/s(泡状流中的滑脱速度Vs平均值可取0.244m/s）
				//Hg = 1 / 2 * (1 + qm / (Vs * At) - sqrt(pow((1 + qm / (Vs * At)), 2) - 4 * qgsc / (Vs * At)));// '空隙率
				Calculate_temp1 = powf((1 + qm/(Vs*At)), 2)
				Calculate_temp2 = 4 * qgsc / (Vs * At)
				Calculate_temp3 = sqrtf(Calculate_temp1 - Calculate_temp2)
				Calculate_temp4 = qm / (Vs * At)
				Calculate_temp5 = 1 + Calculate_temp4 - Calculate_temp3
				Hg = 1 / 2 * Calculate_temp5

				Rhom = (1-Hg)*Rhow + Hg*Rhog
				Re = Rhow * Dti
				Re *= Vsl / Muw
				e = 0.01524 // '绝对粗糙度，mm
				if Re > 2300 {
					//fm = pow((1.14 - 2 * log(e / (DT * pow(10, 3)) + 21.25 / pow(Re, 0.9)) / log(10)), (-2));
					Calculate_temp1 = e / (DT * powf(10, 3))
					Calculate_temp2 = 21.25 / powf(Re, 0.9)
					Calculate_temp3 = 2 * logf(Calculate_temp1+Calculate_temp2)
					Calculate_temp4 = 1.14 - Calculate_temp3/logf(10)
					fm = powf(Calculate_temp4, (-2))
				} else {
					fm = 64 / Re
				}
				//Tauf = fm * Rhow * pow(Vsl, 2) / (2 * Dti * (1 - Hg));
				Calculate_temp1 = fm * Rhow
				Calculate_temp2 = powf(Vsl, 2)
				Calculate_temp3 = 2 * Dti * (1 - Hg)

				Tauf = Calculate_temp1 * Calculate_temp2
				Tauf /= Calculate_temp3

				//gradP = (Rhom * g + Tauf) / (1 - Wm * vm / (At * Pav * pow(10, 6)));
				Calculate_temp1 = Rhom*g + Tauf
				Calculate_temp2 = At * Pav
				Calculate_temp2 = Calculate_temp2 * powf(10, 6)
				Calculate_temp3 = 1 - Wm*vm/Calculate_temp2
				gradP = Calculate_temp1 / Calculate_temp3
				dp0 = gradP * dh * powf(10, -6)
			}
			djsvaue = fabsf(dp0 - dp)

			if (djsvaue / dp) <= 0.00001 {
				break
			}
		}

		H[i] = H[i-1] + dh
		T[i] = T[i-1] + gradT*dh
		gradP1[i] = gradP
		P[i] = P[i-1] + gradP1[i]*dh/powf(10, 6)

	}
}

//输入参数
var (
	//定义井身结构
	Dti float32 = 64   //油管内径，m
	Dto float32 = 73   //油管外径，m
	Dci float32 = 112  //套管内径，m
	Ht  float32 = 2000 //油管下入深度，m
	Hr  float32 = 2100 //储层中深，m
	//定义生产数据
	Pt  float32 = 0        //油压，MPa
	Pc  float32 = 0        //套压，MPa
	Qg  float32 = 1.0      //日产气，万方/天
	Qw  float32 = 1.0      //日产水，方
	Twh float32 = 22 + 273 //井口温度，K
	//定义测试数据
	Pcg float32 = 0
	Tr  float32 = 120 + 273 //地层温度，K
	//
	Yg  float32 = 0.56
	Tav float32 = 0
	Pav float32 = 0
	DT  float32 = 0
	Z   float32 = 0
	//
	Pc0 float32 = 15 //积液前套压,MPa
	Pwf float32 = 0
	Pti float32 = 0
	Ht1 float32 = 0
	//基础数据3
	g float32 = 9.81
	//
	I_Select_JY uint8 = 0
)

//点击计算
// Pt 油压 Pc  套压
//return D_YGJYGD, D_TGJYGD, D_JYL, D_JiaYL
func feedingCalculate(pt, pc float32) (float32, float32, float32, float32, error) {
	//FIXME:需要加上执行超时，避免算法不收敛，导致一直计算，但是以前的这里是没有家超时的，估计不会触发bug
	var djsvaue float32 = 0
	var At float32 = 0
	var Ac float32 = 0
	var htl float32 = 0
	var hcl float32 = 0

	g = 9.81 //重力加速度

	var materialNum modle.MaterialNum
	readMaterialNumCfgTofile(&materialNum)

	Yg = materialNum.Yg
	Dti = materialNum.Dti
	Dto = materialNum.Dto
	Dci = materialNum.Dci
	Ht = materialNum.Ht
	Hr = materialNum.Hr
	Pc0 = materialNum.Pco
	Qg = materialNum.Qg
	Qw = materialNum.Qw
	Twh = materialNum.Twh + 273
	Tr = materialNum.Tr + 273

	Pc = pc
	Pt = pt

	At = 3.14 * (Dti * Dti) / 4         //油管面积，平方米
	Ac = 3.14 / 4 * (Dci*Dci - Dto*Dto) //环空面积，平方米
	//'--------------------------------------------------
	//'1、计算临界携液流量，判断目前工况条件下，是否会积液
	var Rhol float32 = 1050
	var Rhog float32 = f_Rhog(Twh, Pt, Yg)
	var Sigma float32 = f_Sigmaw(Twh, Pt) / 1000
	//	var Mug float32 = f_Mug(Twh, Pt, Yg)
	var Z float32 = z_DPR(Twh, Pt, Yg)
	var dd1 float32 = Sigma * (Rhol - Rhog) / (Rhog * Rhog)
	//float  Vcr = 2.5 * (Sigma * (Rhol - Rhog) / （Rhog*Rhog）)^0.25;// '李闽模型
	var Vcr float32 = 2.5 * (powf(dd1, 0.25))           // '李闽模型
	var qcr float32 = 25000 * At * Pt * Vcr / (Z * Twh) //'临界流量
	var Tav float32
	var Pwf0 float32
	var pwf_1 float32
	var Hcl_0 float32
	var htl_0 float32
	var pwf1 float32
	var Calculate_temp1, Calculate_temp2, Calculate_temp3, Calculate_temp4 float32 //中间变量
	if Qg < qcr {
		//'1.1、由积液前的套压计算井底流压
		DT = (Tr - Twh) / Hr    //  '温度梯度
		Tav = Twh + DT*Hr/2     // '环空气柱的平均温度
		Z = z_DPR(Tav, Pc0, Yg) // '天然气偏差系数
		//Pwf0 = Pc0 * (exp(0.03418 * Yg * Hr / (Tav * Z)))
		Calculate_temp1 = 0.03418 * Yg * Hr
		Calculate_temp2 = Tav * Z
		Pwf0 = Pc0 * (expf(Calculate_temp1 / Calculate_temp2))
		for {
			Pwf = Pwf0              //
			Pav = (Pc0 + Pwf) / 2   //
			Z = z_DPR(Tav, Pav, Yg) //
			//Pwf0 = Pc0 * (exp(0.03418 * Yg * Hr / (Tav * Z)));//
			Calculate_temp1 = 0.03418 * Yg * Hr
			Calculate_temp2 = Tav * Z
			Pwf0 = Pc0 * (expf(Calculate_temp1 / Calculate_temp2))
			djsvaue = fabsf(Pwf - Pwf0)
			if djsvaue <= 0.01 {
				break
			}
		}

		//1.2、计算套管中的积液高度
		//1.2.1、将地层中部深度作为环空气柱的迭代初值
		//pwf_1 = Pc * (exp(0.03418 * Yg * Hr / (Tav * Z)));
		Calculate_temp1 = 0.03418 * Yg * Hr
		Calculate_temp2 = Tav * Z
		pwf_1 = Pc * (expf(Calculate_temp1 / Calculate_temp2))
		if pwf_1 > Pwf {
			return 0, 0, 0, 0, errors.New("请检查积液前后的套压数据是否正确")
		} else {
			Hcl_0 = (Pwf - pwf_1) * 1000000
			Hcl_0 /= (1050 * g)

			for {
				hcl = Hcl_0
				Tav = Twh + DT*(Hr-hcl)/2
				Z = z_DPR(Tav, Pc, Yg)
				//pwf_1 = Pc * (exp(0.03418 * Yg * (Hr - hcl) / (Tav * Z))) + 1050 * g * hcl / 1000000;
				Calculate_temp1 = 0.03418 * Yg * (Hr - hcl)
				Calculate_temp2 = Tav * Z
				Calculate_temp3 = expf(Calculate_temp1 / Calculate_temp2)
				Calculate_temp4 = 1050 * g * hcl
				Calculate_temp4 /= 1000000
				pwf_1 = Pc * Calculate_temp3
				pwf_1 += Calculate_temp4
				if pwf_1 < Pwf {
					Calculate_temp1 = (Pwf - pwf_1) * 1000000
					Hcl_0 = hcl + Calculate_temp1/(1050*g)
				} else {
					Calculate_temp1 = (Pwf - pwf_1) * 1000000
					Hcl_0 = hcl - Calculate_temp1/(1050*g)
				}

				djsvaue = fabsf(pwf_1 - Pwf)
				if djsvaue <= 0.01 {
					break
				}
			}

			//'1.3、计算油管中积液高度
			//'case1：若日产气量为零或者关井状态下,计算油管中积液高度
			if Qg == 0 {
				Tav = (Twh + Tr) / 2    // '环空气柱平均温度，K
				Pav = (Pt + Pwf) / 2    //'环空气柱平均压力，MPa
				Z = z_DPR(Tav, Pav, Yg) // '套管中天然气偏差系数
				//'将地层中部深度作为油管气柱的迭代初值
				//pwf_1 = Pt * (exp(0.03418 * Yg * Hr / (Tav * Z)));//
				Calculate_temp1 = 0.03418 * Yg * Hr
				Calculate_temp2 = Tav * Z
				Calculate_temp3 = expf(Calculate_temp1 / Calculate_temp2)
				pwf_1 = Pt * Calculate_temp3
				if pwf_1 > Pwf {
					return 0, 0, 0, 0, errors.New("请核对基础参数，油套压不匹配，或该井没有积液")
				} else {
					htl_0 = (Pwf - pwf_1) * 1000000
					htl_0 /= (1050 * g)
					for {
						htl = htl_0
						//pwf_1 = Pt * (exp(0.03418 * Yg * (Hr - htl) / (Tav * Z))) + 1050 * g * htl / 1000000;//
						Calculate_temp1 = 0.03418 * Yg * (Hr - hcl)
						Calculate_temp2 = Tav * Z
						Calculate_temp3 = expf(Calculate_temp1 / Calculate_temp2)
						pwf_1 = Pt * Calculate_temp3
						pwf_1 += 1050 * g * htl / 1000000

						if pwf_1 < Pwf {
							Calculate_temp1 = (Pwf - pwf_1) * 1000000
							htl_0 = htl + Calculate_temp1/(1050*g)
						} else {
							Calculate_temp1 = (Pwf - pwf_1) * 1000000
							htl_0 = htl - Calculate_temp1/(1050*g)
						}
						djsvaue = fabsf(pwf_1 - Pwf)
						if djsvaue <= 0.001 {
							break
						}
					}
				}
			} else {
				//Pwh = Pt //
				//qsc = Qg //
				Ht1 = Ht //
				//Call Hagedorn_Brown
				Hagedorn_Brown()
				Calculate_temp1 = 1050 * g * (Hr - Ht)
				pwf1 = Pti + Calculate_temp1/1000000 //
				if pwf1 > Pwf {
					return 0, 0, 0, 0, errors.New("输入数据有误，请检查错误")
				} else {
					htl_0 = (Pwf - pwf1) * 1000000
					htl_0 /= (1050 * g) //
					for {
						htl = htl_0      //
						Ht1 = Ht - htl_0 //
						Hagedorn_Brown()
						htl_0 = (Pwf - Pti) * 1000000
						htl_0 /= (1050 * g) //
						djsvaue = fabsf(htl - htl_0)
						if djsvaue <= 1 {
							break
						}
					}
				}
			}
			//'1.4、数据输出
			D_YGJYGD = htl          //油管积液高度
			D_TGJYGD = hcl          //套管积液高度
			D_JYL = htl*At + hcl*Ac //积液量
			D_JYL = D_JYL / 1000000

			if I_Select_JY == 0 { //加药浓度选择：0 千分之五；1 千分之十
				D_JiaYL = (D_JYL + Qw) * 5 / 0.995 //1000 * 0.005 / 0.995;
			} else {
				D_JiaYL = (D_JYL + Qw) * 10 / 0.99 //1000 * 0.01 / 0.99;
			}
		}
	} else {
		fmt.Println("该井尚无积液")
		D_YGJYGD = 0
		D_TGJYGD = 0
		D_JYL = 0
		D_JiaYL = 0
	}

	return D_YGJYGD, D_TGJYGD, D_JYL, D_JiaYL, nil
}
