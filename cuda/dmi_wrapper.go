package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var adddmi_code cu.Function

type adddmi_args struct {
	arg_Hx unsafe.Pointer
	arg_Hy unsafe.Pointer
	arg_Hz unsafe.Pointer
	arg_mx unsafe.Pointer
	arg_my unsafe.Pointer
	arg_mz unsafe.Pointer
	arg_Dx float32
	arg_Dy float32
	arg_Dz float32
	arg_A  float32
	arg_cx float32
	arg_cy float32
	arg_cz float32
	arg_Nx int
	arg_Ny int
	arg_Nz int
	argptr [16]unsafe.Pointer
}

// Wrapper for adddmi CUDA kernel, asynchronous.
func k_adddmi_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, A float32, cx float32, cy float32, cz float32, Nx int, Ny int, Nz int, cfg *config, str int) {
	if synchronous { // debug
		SyncAll()
	}

	if adddmi_code == 0 {
		adddmi_code = fatbinLoad(adddmi_map, "adddmi")
	}

	var _a_ adddmi_args

	_a_.arg_Hx = Hx
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_Hx)
	_a_.arg_Hy = Hy
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_Hy)
	_a_.arg_Hz = Hz
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Hz)
	_a_.arg_mx = mx
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_mx)
	_a_.arg_my = my
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_my)
	_a_.arg_mz = mz
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_mz)
	_a_.arg_Dx = Dx
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_Dx)
	_a_.arg_Dy = Dy
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_Dy)
	_a_.arg_Dz = Dz
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_Dz)
	_a_.arg_A = A
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_A)
	_a_.arg_cx = cx
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_cx)
	_a_.arg_cy = cy
	_a_.argptr[11] = unsafe.Pointer(&_a_.arg_cy)
	_a_.arg_cz = cz
	_a_.argptr[12] = unsafe.Pointer(&_a_.arg_cz)
	_a_.arg_Nx = Nx
	_a_.argptr[13] = unsafe.Pointer(&_a_.arg_Nx)
	_a_.arg_Ny = Ny
	_a_.argptr[14] = unsafe.Pointer(&_a_.arg_Ny)
	_a_.arg_Nz = Nz
	_a_.argptr[15] = unsafe.Pointer(&_a_.arg_Nz)

	args := _a_.argptr[:]
	cu.LaunchKernel(adddmi_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)

	if synchronous { // debug
		SyncAll()
	}
}

// Wrapper for adddmi CUDA kernel, synchronized.
func k_adddmi(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, A float32, cx float32, cy float32, cz float32, Nx int, Ny int, Nz int, cfg *config) {
	const stream = 0
	k_adddmi_async(Hx, Hy, Hz, mx, my, mz, Dx, Dy, Dz, A, cx, cy, cz, Nx, Ny, Nz, cfg, stream)
	Sync(stream)
}

var adddmi_map = map[int]string{0: "",
	20: adddmi_ptx_20,
	30: adddmi_ptx_30,
	35: adddmi_ptx_35}

const (
	adddmi_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<99>;
	.reg .f32 	%f<153>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f47, [adddmi_param_6];
	ld.param.f32 	%f48, [adddmi_param_7];
	ld.param.f32 	%f49, [adddmi_param_9];
	ld.param.f32 	%f50, [adddmi_param_10];
	ld.param.f32 	%f51, [adddmi_param_11];
	ld.param.u32 	%r5, [adddmi_param_13];
	ld.param.u32 	%r6, [adddmi_param_14];
	ld.param.u32 	%r7, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r8, %r9, %r10;
	.loc 1 18 1
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	.loc 1 19 1
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r6;
	setp.ge.s32	%p2, %r1, %r5;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r3, %r7;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB0_18;

	.loc 1 25 1
	mad.lo.s32 	%r17, %r3, %r6, %r2;
	mad.lo.s32 	%r18, %r17, %r5, %r1;
	mul.wide.s32 	%rd13, %r18, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f6, [%rd19];
	.loc 1 35 1
	setp.gt.s32	%p6, %r1, 0;
	@%p6 bra 	BB0_3;

	mov.f32 	%f144, 0f00000000;
	mov.f32 	%f143, %f144;
	mov.f32 	%f140, %f144;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd21, %r18, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f140, [%rd22+-4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f143, [%rd24+-4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f144, [%rd26+-4];

BB0_4:
	.loc 1 39 1
	mov.f32 	%f11, %f143;
	add.s32 	%r37, %r1, 1;
	setp.lt.s32	%p7, %r37, %r5;
	@%p7 bra 	BB0_6;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB0_7;

BB0_6:
	mul.wide.s32 	%rd28, %r18, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f137, [%rd29+4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f138, [%rd31+4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f139, [%rd33+4];

BB0_7:
	.loc 1 45 1
	add.f32 	%f19, %f49, %f49;
	.loc 2 3608 3
	div.rn.f32 	%f20, %f47, %f19;
	.loc 1 46 1
	mul.f32 	%f58, %f11, %f11;
	fma.rn.f32 	%f59, %f140, %f140, %f58;
	fma.rn.f32 	%f60, %f144, %f144, %f59;
	setp.neu.f32	%p8, %f60, 0f00000000;
	mov.f32 	%f142, %f11;
	@%p8 bra 	BB0_9;

	.loc 1 47 1
	neg.f32 	%f61, %f50;
	mul.f32 	%f62, %f20, %f61;
	mul.f32 	%f63, %f62, %f6;
	sub.f32 	%f140, %f4, %f63;
	.loc 1 49 1
	fma.rn.f32 	%f144, %f62, %f4, %f6;
	mov.f32 	%f142, %f5;

BB0_9:
	.loc 1 51 1
	mov.f32 	%f24, %f142;
	mul.f32 	%f64, %f138, %f138;
	fma.rn.f32 	%f65, %f137, %f137, %f64;
	fma.rn.f32 	%f66, %f139, %f139, %f65;
	setp.eq.f32	%p9, %f66, 0f00000000;
	.loc 1 52 1
	mul.f32 	%f67, %f20, %f50;
	mul.f32 	%f68, %f67, %f6;
	sub.f32 	%f69, %f4, %f68;
	.loc 1 54 1
	fma.rn.f32 	%f70, %f67, %f4, %f6;
	.loc 1 51 1
	selp.f32	%f71, %f69, %f137, %p9;
	selp.f32	%f72, %f5, %f138, %p9;
	selp.f32	%f73, %f70, %f139, %p9;
	.loc 1 57 1
	mul.f32 	%f74, %f50, %f50;
	.loc 2 3608 3
	div.rn.f32 	%f75, %f19, %f74;
	.loc 1 57 31
	sub.f32 	%f76, %f71, %f4;
	sub.f32 	%f77, %f72, %f5;
	sub.f32 	%f78, %f73, %f6;
	sub.f32 	%f79, %f140, %f4;
	add.f32 	%f80, %f79, %f76;
	sub.f32 	%f81, %f24, %f5;
	add.f32 	%f82, %f81, %f77;
	sub.f32 	%f83, %f144, %f6;
	add.f32 	%f84, %f83, %f78;
	.loc 1 57 1
	fma.rn.f32 	%f85, %f75, %f80, %f1;
	fma.rn.f32 	%f26, %f75, %f82, %f2;
	fma.rn.f32 	%f86, %f75, %f84, %f3;
	.loc 1 58 1
	sub.f32 	%f87, %f73, %f144;
	mul.f32 	%f88, %f87, %f47;
	.loc 2 3608 3
	div.rn.f32 	%f89, %f88, %f50;
	.loc 1 58 55
	add.f32 	%f27, %f85, %f89;
	.loc 1 59 1
	sub.f32 	%f90, %f71, %f140;
	mul.f32 	%f91, %f90, %f47;
	.loc 2 3608 3
	div.rn.f32 	%f92, %f91, %f50;
	.loc 1 59 55
	sub.f32 	%f28, %f86, %f92;
	.loc 1 68 1
	setp.gt.s32	%p10, %r2, 0;
	@%p10 bra 	BB0_11;

	mov.f32 	%f152, 0f00000000;
	mov.f32 	%f151, %f152;
	mov.f32 	%f150, %f152;
	bra.uni 	BB0_12;

BB0_11:
	.loc 1 69 1
	add.s32 	%r65, %r17, -1;
	mad.lo.s32 	%r70, %r65, %r5, %r1;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f150, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f151, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f152, [%rd40];

BB0_12:
	.loc 1 72 1
	mov.f32 	%f32, %f150;
	add.s32 	%r4, %r2, 1;
	setp.lt.s32	%p11, %r4, %r6;
	@%p11 bra 	BB0_14;

	mov.f32 	%f147, 0f00000000;
	mov.f32 	%f146, %f147;
	mov.f32 	%f145, %f147;
	bra.uni 	BB0_15;

BB0_14:
	.loc 1 73 1
	mad.lo.s32 	%r79, %r3, %r6, %r4;
	mad.lo.s32 	%r84, %r79, %r5, %r1;
	mul.wide.s32 	%rd42, %r84, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f145, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f146, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f147, [%rd47];

BB0_15:
	.loc 2 3608 3
	div.rn.f32 	%f41, %f48, %f19;
	.loc 1 78 1
	mul.f32 	%f99, %f151, %f151;
	fma.rn.f32 	%f100, %f32, %f32, %f99;
	fma.rn.f32 	%f101, %f152, %f152, %f100;
	setp.neu.f32	%p12, %f101, 0f00000000;
	mov.f32 	%f149, %f32;
	@%p12 bra 	BB0_17;

	.loc 1 80 1
	neg.f32 	%f102, %f51;
	mul.f32 	%f103, %f41, %f102;
	mul.f32 	%f104, %f103, %f6;
	sub.f32 	%f151, %f5, %f104;
	.loc 1 81 1
	fma.rn.f32 	%f152, %f103, %f5, %f6;
	mov.f32 	%f149, %f4;

BB0_17:
	.loc 1 83 1
	mov.f32 	%f44, %f149;
	mul.f32 	%f105, %f146, %f146;
	fma.rn.f32 	%f106, %f145, %f145, %f105;
	fma.rn.f32 	%f107, %f147, %f147, %f106;
	setp.eq.f32	%p13, %f107, 0f00000000;
	.loc 1 85 1
	mul.f32 	%f108, %f41, %f51;
	mul.f32 	%f109, %f108, %f6;
	sub.f32 	%f110, %f5, %f109;
	.loc 1 86 1
	fma.rn.f32 	%f111, %f108, %f5, %f6;
	.loc 1 83 1
	selp.f32	%f112, %f4, %f145, %p13;
	selp.f32	%f113, %f110, %f146, %p13;
	selp.f32	%f114, %f111, %f147, %p13;
	.loc 1 89 1
	mul.f32 	%f115, %f51, %f51;
	.loc 2 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 89 31
	sub.f32 	%f117, %f112, %f4;
	sub.f32 	%f118, %f113, %f5;
	sub.f32 	%f119, %f114, %f6;
	sub.f32 	%f120, %f44, %f4;
	add.f32 	%f121, %f120, %f117;
	sub.f32 	%f122, %f151, %f5;
	add.f32 	%f123, %f122, %f118;
	sub.f32 	%f124, %f152, %f6;
	add.f32 	%f125, %f124, %f119;
	.loc 1 89 1
	fma.rn.f32 	%f126, %f116, %f121, %f27;
	fma.rn.f32 	%f127, %f116, %f123, %f26;
	fma.rn.f32 	%f128, %f116, %f125, %f28;
	.loc 1 90 1
	sub.f32 	%f129, %f114, %f152;
	mul.f32 	%f130, %f129, %f48;
	.loc 2 3608 3
	div.rn.f32 	%f131, %f130, %f51;
	.loc 1 90 55
	add.f32 	%f132, %f127, %f131;
	.loc 1 91 1
	sub.f32 	%f133, %f113, %f151;
	mul.f32 	%f134, %f133, %f48;
	.loc 2 3608 3
	div.rn.f32 	%f135, %f134, %f51;
	.loc 1 91 55
	sub.f32 	%f136, %f128, %f135;
	mul.wide.s32 	%rd49, %r18, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 95 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd52], %f132;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd54], %f136;

BB0_18:
	.loc 1 98 2
	ret;
}


`
	adddmi_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<99>;
	.reg .f32 	%f<153>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f47, [adddmi_param_6];
	ld.param.f32 	%f48, [adddmi_param_7];
	ld.param.f32 	%f49, [adddmi_param_9];
	ld.param.f32 	%f50, [adddmi_param_10];
	ld.param.f32 	%f51, [adddmi_param_11];
	ld.param.u32 	%r5, [adddmi_param_13];
	ld.param.u32 	%r6, [adddmi_param_14];
	ld.param.u32 	%r7, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r8, %r9, %r10;
	.loc 1 18 1
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	.loc 1 19 1
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r6;
	setp.ge.s32	%p2, %r1, %r5;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r3, %r7;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB0_18;

	.loc 1 25 1
	mad.lo.s32 	%r17, %r3, %r6, %r2;
	mad.lo.s32 	%r18, %r17, %r5, %r1;
	mul.wide.s32 	%rd13, %r18, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f6, [%rd19];
	.loc 1 35 1
	setp.gt.s32	%p6, %r1, 0;
	@%p6 bra 	BB0_3;

	mov.f32 	%f144, 0f00000000;
	mov.f32 	%f143, %f144;
	mov.f32 	%f140, %f144;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd21, %r18, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f140, [%rd22+-4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f143, [%rd24+-4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 37 1
	ld.global.f32 	%f144, [%rd26+-4];

BB0_4:
	.loc 1 39 1
	mov.f32 	%f11, %f143;
	add.s32 	%r37, %r1, 1;
	setp.lt.s32	%p7, %r37, %r5;
	@%p7 bra 	BB0_6;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB0_7;

BB0_6:
	mul.wide.s32 	%rd28, %r18, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f137, [%rd29+4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f138, [%rd31+4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 41 1
	ld.global.f32 	%f139, [%rd33+4];

BB0_7:
	.loc 1 45 1
	add.f32 	%f19, %f49, %f49;
	.loc 2 3608 3
	div.rn.f32 	%f20, %f47, %f19;
	.loc 1 46 1
	mul.f32 	%f58, %f11, %f11;
	fma.rn.f32 	%f59, %f140, %f140, %f58;
	fma.rn.f32 	%f60, %f144, %f144, %f59;
	setp.neu.f32	%p8, %f60, 0f00000000;
	mov.f32 	%f142, %f11;
	@%p8 bra 	BB0_9;

	.loc 1 47 1
	neg.f32 	%f61, %f50;
	mul.f32 	%f62, %f20, %f61;
	mul.f32 	%f63, %f62, %f6;
	sub.f32 	%f140, %f4, %f63;
	.loc 1 49 1
	fma.rn.f32 	%f144, %f62, %f4, %f6;
	mov.f32 	%f142, %f5;

BB0_9:
	.loc 1 51 1
	mov.f32 	%f24, %f142;
	mul.f32 	%f64, %f138, %f138;
	fma.rn.f32 	%f65, %f137, %f137, %f64;
	fma.rn.f32 	%f66, %f139, %f139, %f65;
	setp.eq.f32	%p9, %f66, 0f00000000;
	.loc 1 52 1
	mul.f32 	%f67, %f20, %f50;
	mul.f32 	%f68, %f67, %f6;
	sub.f32 	%f69, %f4, %f68;
	.loc 1 54 1
	fma.rn.f32 	%f70, %f67, %f4, %f6;
	.loc 1 51 1
	selp.f32	%f71, %f69, %f137, %p9;
	selp.f32	%f72, %f5, %f138, %p9;
	selp.f32	%f73, %f70, %f139, %p9;
	.loc 1 57 1
	mul.f32 	%f74, %f50, %f50;
	.loc 2 3608 3
	div.rn.f32 	%f75, %f19, %f74;
	.loc 1 57 31
	sub.f32 	%f76, %f71, %f4;
	sub.f32 	%f77, %f72, %f5;
	sub.f32 	%f78, %f73, %f6;
	sub.f32 	%f79, %f140, %f4;
	add.f32 	%f80, %f79, %f76;
	sub.f32 	%f81, %f24, %f5;
	add.f32 	%f82, %f81, %f77;
	sub.f32 	%f83, %f144, %f6;
	add.f32 	%f84, %f83, %f78;
	.loc 1 57 1
	fma.rn.f32 	%f85, %f75, %f80, %f1;
	fma.rn.f32 	%f26, %f75, %f82, %f2;
	fma.rn.f32 	%f86, %f75, %f84, %f3;
	.loc 1 58 1
	sub.f32 	%f87, %f73, %f144;
	mul.f32 	%f88, %f87, %f47;
	.loc 2 3608 3
	div.rn.f32 	%f89, %f88, %f50;
	.loc 1 58 55
	add.f32 	%f27, %f85, %f89;
	.loc 1 59 1
	sub.f32 	%f90, %f71, %f140;
	mul.f32 	%f91, %f90, %f47;
	.loc 2 3608 3
	div.rn.f32 	%f92, %f91, %f50;
	.loc 1 59 55
	sub.f32 	%f28, %f86, %f92;
	.loc 1 68 1
	setp.gt.s32	%p10, %r2, 0;
	@%p10 bra 	BB0_11;

	mov.f32 	%f152, 0f00000000;
	mov.f32 	%f151, %f152;
	mov.f32 	%f150, %f152;
	bra.uni 	BB0_12;

BB0_11:
	.loc 1 69 1
	add.s32 	%r65, %r17, -1;
	mad.lo.s32 	%r70, %r65, %r5, %r1;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f150, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f151, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 70 1
	ld.global.f32 	%f152, [%rd40];

BB0_12:
	.loc 1 72 1
	mov.f32 	%f32, %f150;
	add.s32 	%r4, %r2, 1;
	setp.lt.s32	%p11, %r4, %r6;
	@%p11 bra 	BB0_14;

	mov.f32 	%f147, 0f00000000;
	mov.f32 	%f146, %f147;
	mov.f32 	%f145, %f147;
	bra.uni 	BB0_15;

BB0_14:
	.loc 1 73 1
	mad.lo.s32 	%r79, %r3, %r6, %r4;
	mad.lo.s32 	%r84, %r79, %r5, %r1;
	mul.wide.s32 	%rd42, %r84, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f145, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f146, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 74 1
	ld.global.f32 	%f147, [%rd47];

BB0_15:
	.loc 2 3608 3
	div.rn.f32 	%f41, %f48, %f19;
	.loc 1 78 1
	mul.f32 	%f99, %f151, %f151;
	fma.rn.f32 	%f100, %f32, %f32, %f99;
	fma.rn.f32 	%f101, %f152, %f152, %f100;
	setp.neu.f32	%p12, %f101, 0f00000000;
	mov.f32 	%f149, %f32;
	@%p12 bra 	BB0_17;

	.loc 1 80 1
	neg.f32 	%f102, %f51;
	mul.f32 	%f103, %f41, %f102;
	mul.f32 	%f104, %f103, %f6;
	sub.f32 	%f151, %f5, %f104;
	.loc 1 81 1
	fma.rn.f32 	%f152, %f103, %f5, %f6;
	mov.f32 	%f149, %f4;

BB0_17:
	.loc 1 83 1
	mov.f32 	%f44, %f149;
	mul.f32 	%f105, %f146, %f146;
	fma.rn.f32 	%f106, %f145, %f145, %f105;
	fma.rn.f32 	%f107, %f147, %f147, %f106;
	setp.eq.f32	%p13, %f107, 0f00000000;
	.loc 1 85 1
	mul.f32 	%f108, %f41, %f51;
	mul.f32 	%f109, %f108, %f6;
	sub.f32 	%f110, %f5, %f109;
	.loc 1 86 1
	fma.rn.f32 	%f111, %f108, %f5, %f6;
	.loc 1 83 1
	selp.f32	%f112, %f4, %f145, %p13;
	selp.f32	%f113, %f110, %f146, %p13;
	selp.f32	%f114, %f111, %f147, %p13;
	.loc 1 89 1
	mul.f32 	%f115, %f51, %f51;
	.loc 2 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 89 31
	sub.f32 	%f117, %f112, %f4;
	sub.f32 	%f118, %f113, %f5;
	sub.f32 	%f119, %f114, %f6;
	sub.f32 	%f120, %f44, %f4;
	add.f32 	%f121, %f120, %f117;
	sub.f32 	%f122, %f151, %f5;
	add.f32 	%f123, %f122, %f118;
	sub.f32 	%f124, %f152, %f6;
	add.f32 	%f125, %f124, %f119;
	.loc 1 89 1
	fma.rn.f32 	%f126, %f116, %f121, %f27;
	fma.rn.f32 	%f127, %f116, %f123, %f26;
	fma.rn.f32 	%f128, %f116, %f125, %f28;
	.loc 1 90 1
	sub.f32 	%f129, %f114, %f152;
	mul.f32 	%f130, %f129, %f48;
	.loc 2 3608 3
	div.rn.f32 	%f131, %f130, %f51;
	.loc 1 90 55
	add.f32 	%f132, %f127, %f131;
	.loc 1 91 1
	sub.f32 	%f133, %f113, %f151;
	mul.f32 	%f134, %f133, %f48;
	.loc 2 3608 3
	div.rn.f32 	%f135, %f134, %f51;
	.loc 1 91 55
	sub.f32 	%f136, %f128, %f135;
	mul.wide.s32 	%rd49, %r18, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 95 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd52], %f132;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd54], %f136;

BB0_18:
	.loc 1 98 2
	ret;
}


`
	adddmi_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<99>;
	.reg .f32 	%f<153>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f47, [adddmi_param_6];
	ld.param.f32 	%f48, [adddmi_param_7];
	ld.param.f32 	%f49, [adddmi_param_9];
	ld.param.f32 	%f50, [adddmi_param_10];
	ld.param.f32 	%f51, [adddmi_param_11];
	ld.param.u32 	%r5, [adddmi_param_13];
	ld.param.u32 	%r6, [adddmi_param_14];
	ld.param.u32 	%r7, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r8, %r9, %r10;
	.loc 1 18 1
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	.loc 1 19 1
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r6;
	setp.ge.s32	%p2, %r1, %r5;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r3, %r7;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB2_18;

	.loc 1 25 1
	mad.lo.s32 	%r17, %r3, %r6, %r2;
	mad.lo.s32 	%r18, %r17, %r5, %r1;
	mul.wide.s32 	%rd13, %r18, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f6, [%rd19];
	.loc 1 35 1
	setp.gt.s32	%p6, %r1, 0;
	@%p6 bra 	BB2_3;

	mov.f32 	%f144, 0f00000000;
	mov.f32 	%f143, %f144;
	mov.f32 	%f140, %f144;
	bra.uni 	BB2_4;

BB2_3:
	mul.wide.s32 	%rd21, %r18, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 37 1
	ld.global.nc.f32 	%f140, [%rd22+-4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 37 1
	ld.global.nc.f32 	%f143, [%rd24+-4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 37 1
	ld.global.nc.f32 	%f144, [%rd26+-4];

BB2_4:
	.loc 1 39 1
	mov.f32 	%f11, %f143;
	add.s32 	%r37, %r1, 1;
	setp.lt.s32	%p7, %r37, %r5;
	@%p7 bra 	BB2_6;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB2_7;

BB2_6:
	mul.wide.s32 	%rd28, %r18, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 41 1
	ld.global.nc.f32 	%f137, [%rd29+4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 41 1
	ld.global.nc.f32 	%f138, [%rd31+4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 41 1
	ld.global.nc.f32 	%f139, [%rd33+4];

BB2_7:
	.loc 1 45 1
	add.f32 	%f19, %f49, %f49;
	.loc 3 3608 3
	div.rn.f32 	%f20, %f47, %f19;
	.loc 1 46 1
	mul.f32 	%f58, %f11, %f11;
	fma.rn.f32 	%f59, %f140, %f140, %f58;
	fma.rn.f32 	%f60, %f144, %f144, %f59;
	setp.neu.f32	%p8, %f60, 0f00000000;
	mov.f32 	%f142, %f11;
	@%p8 bra 	BB2_9;

	.loc 1 47 1
	neg.f32 	%f61, %f50;
	mul.f32 	%f62, %f20, %f61;
	mul.f32 	%f63, %f62, %f6;
	sub.f32 	%f140, %f4, %f63;
	.loc 1 49 1
	fma.rn.f32 	%f144, %f62, %f4, %f6;
	mov.f32 	%f142, %f5;

BB2_9:
	.loc 1 51 1
	mov.f32 	%f24, %f142;
	mul.f32 	%f64, %f138, %f138;
	fma.rn.f32 	%f65, %f137, %f137, %f64;
	fma.rn.f32 	%f66, %f139, %f139, %f65;
	setp.eq.f32	%p9, %f66, 0f00000000;
	.loc 1 52 1
	mul.f32 	%f67, %f20, %f50;
	mul.f32 	%f68, %f67, %f6;
	sub.f32 	%f69, %f4, %f68;
	.loc 1 54 1
	fma.rn.f32 	%f70, %f67, %f4, %f6;
	.loc 1 51 1
	selp.f32	%f71, %f69, %f137, %p9;
	selp.f32	%f72, %f5, %f138, %p9;
	selp.f32	%f73, %f70, %f139, %p9;
	.loc 1 57 1
	mul.f32 	%f74, %f50, %f50;
	.loc 3 3608 3
	div.rn.f32 	%f75, %f19, %f74;
	.loc 1 57 31
	sub.f32 	%f76, %f71, %f4;
	sub.f32 	%f77, %f72, %f5;
	sub.f32 	%f78, %f73, %f6;
	sub.f32 	%f79, %f140, %f4;
	add.f32 	%f80, %f79, %f76;
	sub.f32 	%f81, %f24, %f5;
	add.f32 	%f82, %f81, %f77;
	sub.f32 	%f83, %f144, %f6;
	add.f32 	%f84, %f83, %f78;
	.loc 1 57 1
	fma.rn.f32 	%f85, %f75, %f80, %f1;
	fma.rn.f32 	%f26, %f75, %f82, %f2;
	fma.rn.f32 	%f86, %f75, %f84, %f3;
	.loc 1 58 1
	sub.f32 	%f87, %f73, %f144;
	mul.f32 	%f88, %f87, %f47;
	.loc 3 3608 3
	div.rn.f32 	%f89, %f88, %f50;
	.loc 1 58 55
	add.f32 	%f27, %f85, %f89;
	.loc 1 59 1
	sub.f32 	%f90, %f71, %f140;
	mul.f32 	%f91, %f90, %f47;
	.loc 3 3608 3
	div.rn.f32 	%f92, %f91, %f50;
	.loc 1 59 55
	sub.f32 	%f28, %f86, %f92;
	.loc 1 68 1
	setp.gt.s32	%p10, %r2, 0;
	@%p10 bra 	BB2_11;

	mov.f32 	%f152, 0f00000000;
	mov.f32 	%f151, %f152;
	mov.f32 	%f150, %f152;
	bra.uni 	BB2_12;

BB2_11:
	.loc 1 69 1
	add.s32 	%r65, %r17, -1;
	mad.lo.s32 	%r70, %r65, %r5, %r1;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 70 1
	ld.global.nc.f32 	%f150, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 70 1
	ld.global.nc.f32 	%f151, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 70 1
	ld.global.nc.f32 	%f152, [%rd40];

BB2_12:
	.loc 1 72 1
	mov.f32 	%f32, %f150;
	add.s32 	%r4, %r2, 1;
	setp.lt.s32	%p11, %r4, %r6;
	@%p11 bra 	BB2_14;

	mov.f32 	%f147, 0f00000000;
	mov.f32 	%f146, %f147;
	mov.f32 	%f145, %f147;
	bra.uni 	BB2_15;

BB2_14:
	.loc 1 73 1
	mad.lo.s32 	%r79, %r3, %r6, %r4;
	mad.lo.s32 	%r84, %r79, %r5, %r1;
	mul.wide.s32 	%rd42, %r84, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 74 1
	ld.global.nc.f32 	%f145, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 74 1
	ld.global.nc.f32 	%f146, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 74 1
	ld.global.nc.f32 	%f147, [%rd47];

BB2_15:
	.loc 3 3608 3
	div.rn.f32 	%f41, %f48, %f19;
	.loc 1 78 1
	mul.f32 	%f99, %f151, %f151;
	fma.rn.f32 	%f100, %f32, %f32, %f99;
	fma.rn.f32 	%f101, %f152, %f152, %f100;
	setp.neu.f32	%p12, %f101, 0f00000000;
	mov.f32 	%f149, %f32;
	@%p12 bra 	BB2_17;

	.loc 1 80 1
	neg.f32 	%f102, %f51;
	mul.f32 	%f103, %f41, %f102;
	mul.f32 	%f104, %f103, %f6;
	sub.f32 	%f151, %f5, %f104;
	.loc 1 81 1
	fma.rn.f32 	%f152, %f103, %f5, %f6;
	mov.f32 	%f149, %f4;

BB2_17:
	.loc 1 83 1
	mov.f32 	%f44, %f149;
	mul.f32 	%f105, %f146, %f146;
	fma.rn.f32 	%f106, %f145, %f145, %f105;
	fma.rn.f32 	%f107, %f147, %f147, %f106;
	setp.eq.f32	%p13, %f107, 0f00000000;
	.loc 1 85 1
	mul.f32 	%f108, %f41, %f51;
	mul.f32 	%f109, %f108, %f6;
	sub.f32 	%f110, %f5, %f109;
	.loc 1 86 1
	fma.rn.f32 	%f111, %f108, %f5, %f6;
	.loc 1 83 1
	selp.f32	%f112, %f4, %f145, %p13;
	selp.f32	%f113, %f110, %f146, %p13;
	selp.f32	%f114, %f111, %f147, %p13;
	.loc 1 89 1
	mul.f32 	%f115, %f51, %f51;
	.loc 3 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 89 31
	sub.f32 	%f117, %f112, %f4;
	sub.f32 	%f118, %f113, %f5;
	sub.f32 	%f119, %f114, %f6;
	sub.f32 	%f120, %f44, %f4;
	add.f32 	%f121, %f120, %f117;
	sub.f32 	%f122, %f151, %f5;
	add.f32 	%f123, %f122, %f118;
	sub.f32 	%f124, %f152, %f6;
	add.f32 	%f125, %f124, %f119;
	.loc 1 89 1
	fma.rn.f32 	%f126, %f116, %f121, %f27;
	fma.rn.f32 	%f127, %f116, %f123, %f26;
	fma.rn.f32 	%f128, %f116, %f125, %f28;
	.loc 1 90 1
	sub.f32 	%f129, %f114, %f152;
	mul.f32 	%f130, %f129, %f48;
	.loc 3 3608 3
	div.rn.f32 	%f131, %f130, %f51;
	.loc 1 90 55
	add.f32 	%f132, %f127, %f131;
	.loc 1 91 1
	sub.f32 	%f133, %f113, %f151;
	mul.f32 	%f134, %f133, %f48;
	.loc 3 3608 3
	div.rn.f32 	%f135, %f134, %f51;
	.loc 1 91 55
	sub.f32 	%f136, %f128, %f135;
	mul.wide.s32 	%rd49, %r18, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 95 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd52], %f132;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd54], %f136;

BB2_18:
	.loc 1 98 2
	ret;
}


`
)
