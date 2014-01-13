package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for addslonczewskitorque kernel
var addslonczewskitorque_code cu.Function

// Stores the arguments for addslonczewskitorque kernel invocation
type addslonczewskitorque_args_t struct {
	arg_tx              unsafe.Pointer
	arg_ty              unsafe.Pointer
	arg_tz              unsafe.Pointer
	arg_mx              unsafe.Pointer
	arg_my              unsafe.Pointer
	arg_mz              unsafe.Pointer
	arg_jz              unsafe.Pointer
	arg_pxLUT           unsafe.Pointer
	arg_pyLUT           unsafe.Pointer
	arg_pzLUT           unsafe.Pointer
	arg_msatLUT         unsafe.Pointer
	arg_alphaLUT        unsafe.Pointer
	arg_flt             float32
	arg_polLUT          unsafe.Pointer
	arg_lambdaLUT       unsafe.Pointer
	arg_epsilonPrimeLUT unsafe.Pointer
	arg_regions         unsafe.Pointer
	arg_N               int
	argptr              [18]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for addslonczewskitorque kernel invocation
var addslonczewskitorque_args addslonczewskitorque_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	addslonczewskitorque_args.argptr[0] = unsafe.Pointer(&addslonczewskitorque_args.arg_tx)
	addslonczewskitorque_args.argptr[1] = unsafe.Pointer(&addslonczewskitorque_args.arg_ty)
	addslonczewskitorque_args.argptr[2] = unsafe.Pointer(&addslonczewskitorque_args.arg_tz)
	addslonczewskitorque_args.argptr[3] = unsafe.Pointer(&addslonczewskitorque_args.arg_mx)
	addslonczewskitorque_args.argptr[4] = unsafe.Pointer(&addslonczewskitorque_args.arg_my)
	addslonczewskitorque_args.argptr[5] = unsafe.Pointer(&addslonczewskitorque_args.arg_mz)
	addslonczewskitorque_args.argptr[6] = unsafe.Pointer(&addslonczewskitorque_args.arg_jz)
	addslonczewskitorque_args.argptr[7] = unsafe.Pointer(&addslonczewskitorque_args.arg_pxLUT)
	addslonczewskitorque_args.argptr[8] = unsafe.Pointer(&addslonczewskitorque_args.arg_pyLUT)
	addslonczewskitorque_args.argptr[9] = unsafe.Pointer(&addslonczewskitorque_args.arg_pzLUT)
	addslonczewskitorque_args.argptr[10] = unsafe.Pointer(&addslonczewskitorque_args.arg_msatLUT)
	addslonczewskitorque_args.argptr[11] = unsafe.Pointer(&addslonczewskitorque_args.arg_alphaLUT)
	addslonczewskitorque_args.argptr[12] = unsafe.Pointer(&addslonczewskitorque_args.arg_flt)
	addslonczewskitorque_args.argptr[13] = unsafe.Pointer(&addslonczewskitorque_args.arg_polLUT)
	addslonczewskitorque_args.argptr[14] = unsafe.Pointer(&addslonczewskitorque_args.arg_lambdaLUT)
	addslonczewskitorque_args.argptr[15] = unsafe.Pointer(&addslonczewskitorque_args.arg_epsilonPrimeLUT)
	addslonczewskitorque_args.argptr[16] = unsafe.Pointer(&addslonczewskitorque_args.arg_regions)
	addslonczewskitorque_args.argptr[17] = unsafe.Pointer(&addslonczewskitorque_args.arg_N)
}

// Wrapper for addslonczewskitorque CUDA kernel, asynchronous.
func k_addslonczewskitorque_async(tx unsafe.Pointer, ty unsafe.Pointer, tz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, jz unsafe.Pointer, pxLUT unsafe.Pointer, pyLUT unsafe.Pointer, pzLUT unsafe.Pointer, msatLUT unsafe.Pointer, alphaLUT unsafe.Pointer, flt float32, polLUT unsafe.Pointer, lambdaLUT unsafe.Pointer, epsilonPrimeLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	addslonczewskitorque_args.Lock()
	defer addslonczewskitorque_args.Unlock()

	if addslonczewskitorque_code == 0 {
		addslonczewskitorque_code = fatbinLoad(addslonczewskitorque_map, "addslonczewskitorque")
	}

	addslonczewskitorque_args.arg_tx = tx
	addslonczewskitorque_args.arg_ty = ty
	addslonczewskitorque_args.arg_tz = tz
	addslonczewskitorque_args.arg_mx = mx
	addslonczewskitorque_args.arg_my = my
	addslonczewskitorque_args.arg_mz = mz
	addslonczewskitorque_args.arg_jz = jz
	addslonczewskitorque_args.arg_pxLUT = pxLUT
	addslonczewskitorque_args.arg_pyLUT = pyLUT
	addslonczewskitorque_args.arg_pzLUT = pzLUT
	addslonczewskitorque_args.arg_msatLUT = msatLUT
	addslonczewskitorque_args.arg_alphaLUT = alphaLUT
	addslonczewskitorque_args.arg_flt = flt
	addslonczewskitorque_args.arg_polLUT = polLUT
	addslonczewskitorque_args.arg_lambdaLUT = lambdaLUT
	addslonczewskitorque_args.arg_epsilonPrimeLUT = epsilonPrimeLUT
	addslonczewskitorque_args.arg_regions = regions
	addslonczewskitorque_args.arg_N = N

	args := addslonczewskitorque_args.argptr[:]
	cu.LaunchKernel(addslonczewskitorque_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for addslonczewskitorque kernel.
var addslonczewskitorque_map = map[int]string{0: "",
	20: addslonczewskitorque_ptx_20,
	30: addslonczewskitorque_ptx_30,
	35: addslonczewskitorque_ptx_35}

// addslonczewskitorque PTX code for various compute capabilities.
const (
	addslonczewskitorque_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry addslonczewskitorque(
	.param .u64 addslonczewskitorque_param_0,
	.param .u64 addslonczewskitorque_param_1,
	.param .u64 addslonczewskitorque_param_2,
	.param .u64 addslonczewskitorque_param_3,
	.param .u64 addslonczewskitorque_param_4,
	.param .u64 addslonczewskitorque_param_5,
	.param .u64 addslonczewskitorque_param_6,
	.param .u64 addslonczewskitorque_param_7,
	.param .u64 addslonczewskitorque_param_8,
	.param .u64 addslonczewskitorque_param_9,
	.param .u64 addslonczewskitorque_param_10,
	.param .u64 addslonczewskitorque_param_11,
	.param .f32 addslonczewskitorque_param_12,
	.param .u64 addslonczewskitorque_param_13,
	.param .u64 addslonczewskitorque_param_14,
	.param .u64 addslonczewskitorque_param_15,
	.param .u64 addslonczewskitorque_param_16,
	.param .u32 addslonczewskitorque_param_17
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<35>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<55>;
	.reg .f64 	%fd<3>;


	ld.param.u64 	%rd3, [addslonczewskitorque_param_0];
	ld.param.u64 	%rd4, [addslonczewskitorque_param_1];
	ld.param.u64 	%rd5, [addslonczewskitorque_param_2];
	ld.param.u64 	%rd6, [addslonczewskitorque_param_3];
	ld.param.u64 	%rd7, [addslonczewskitorque_param_4];
	ld.param.u64 	%rd8, [addslonczewskitorque_param_5];
	ld.param.u64 	%rd9, [addslonczewskitorque_param_6];
	ld.param.u64 	%rd10, [addslonczewskitorque_param_7];
	ld.param.u64 	%rd11, [addslonczewskitorque_param_8];
	ld.param.u64 	%rd12, [addslonczewskitorque_param_9];
	ld.param.u64 	%rd13, [addslonczewskitorque_param_10];
	ld.param.u64 	%rd14, [addslonczewskitorque_param_11];
	ld.param.f32 	%f19, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd18, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	cvta.to.global.u64 	%rd1, %rd18;
	.loc 2 16 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 17 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_5;

	cvta.to.global.u64 	%rd19, %rd6;
	.loc 2 19 1
	cvt.s64.s32 	%rd20, %r1;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd19, %rd21;
	ld.global.f32 	%f1, [%rd22];
	cvta.to.global.u64 	%rd23, %rd7;
	.loc 2 19 1
	add.s64 	%rd24, %rd23, %rd21;
	ld.global.f32 	%f2, [%rd24];
	cvta.to.global.u64 	%rd25, %rd8;
	.loc 2 19 1
	add.s64 	%rd26, %rd25, %rd21;
	ld.global.f32 	%f3, [%rd26];
	cvta.to.global.u64 	%rd27, %rd9;
	.loc 2 20 1
	add.s64 	%rd28, %rd27, %rd21;
	ld.global.f32 	%f4, [%rd28];
	.loc 2 23 1
	add.s64 	%rd29, %rd1, %rd20;
	.loc 2 25 1
	ld.global.u8 	%rd2, [%rd29];
	cvta.to.global.u64 	%rd30, %rd10;
	.loc 2 25 1
	shl.b64 	%rd31, %rd2, 2;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd11;
	.loc 2 25 1
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd12;
	.loc 2 25 1
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.f32 	%f5, [%rd32];
	ld.global.f32 	%f6, [%rd34];
	mul.f32 	%f21, %f6, %f6;
	fma.rn.f32 	%f22, %f5, %f5, %f21;
	ld.global.f32 	%f7, [%rd36];
	fma.rn.f32 	%f23, %f7, %f7, %f22;
	.loc 3 991 5
	sqrt.rn.f32 	%f8, %f23;
	mov.f32 	%f76, 0f00000000;
	.loc 2 25 1
	setp.eq.f32 	%p2, %f8, 0f00000000;
	@%p2 bra 	BB0_3;

	rcp.rn.f32 	%f76, %f8;

BB0_3:
	mul.f32 	%f11, %f76, %f5;
	mul.f32 	%f12, %f76, %f6;
	mul.f32 	%f13, %f76, %f7;
	cvta.to.global.u64 	%rd37, %rd13;
	.loc 2 26 1
	add.s64 	%rd39, %rd37, %rd31;
	cvta.to.global.u64 	%rd40, %rd14;
	.loc 2 27 1
	add.s64 	%rd41, %rd40, %rd31;
	ld.global.f32 	%f14, [%rd41];
	cvta.to.global.u64 	%rd42, %rd15;
	.loc 2 28 1
	add.s64 	%rd43, %rd42, %rd31;
	ld.global.f32 	%f15, [%rd43];
	cvta.to.global.u64 	%rd44, %rd16;
	.loc 2 29 1
	add.s64 	%rd45, %rd44, %rd31;
	ld.global.f32 	%f16, [%rd45];
	cvta.to.global.u64 	%rd46, %rd17;
	.loc 2 30 1
	add.s64 	%rd47, %rd46, %rd31;
	ld.global.f32 	%f17, [%rd47];
	.loc 2 26 1
	ld.global.f32 	%f18, [%rd39];
	.loc 2 32 1
	setp.eq.f32 	%p3, %f18, 0f00000000;
	setp.eq.f32 	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB0_5;

	.loc 2 36 1
	mul.f32 	%f24, %f18, %f19;
	.loc 4 2399 3
	div.rn.f32 	%f25, %f4, %f24;
	.loc 2 36 1
	cvt.f64.f32 	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64 	%f26, %fd2;
	.loc 2 37 1
	mul.f32 	%f27, %f16, %f16;
	.loc 2 38 1
	mul.f32 	%f28, %f15, %f27;
	fma.rn.f32 	%f29, %f16, %f16, 0f3F800000;
	fma.rn.f32 	%f30, %f16, %f16, 0fBF800000;
	mul.f32 	%f31, %f12, %f2;
	fma.rn.f32 	%f32, %f11, %f1, %f31;
	fma.rn.f32 	%f33, %f13, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	.loc 4 2399 3
	div.rn.f32 	%f35, %f28, %f34;
	.loc 2 40 1
	mul.f32 	%f36, %f26, %f35;
	.loc 2 41 1
	mul.f32 	%f37, %f26, %f17;
	.loc 2 43 1
	fma.rn.f32 	%f38, %f14, %f14, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	.loc 2 44 1
	mul.f32 	%f40, %f14, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	.loc 2 45 1
	mul.f32 	%f43, %f14, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	.loc 2 47 1
	mul.f32 	%f46, %f13, %f2;
	mul.f32 	%f47, %f12, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f11, %f3;
	mul.f32 	%f50, %f13, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f12, %f1;
	mul.f32 	%f53, %f11, %f2;
	sub.f32 	%f54, %f53, %f52;
	.loc 2 48 1
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	.loc 2 50 1
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	cvta.to.global.u64 	%rd48, %rd3;
	.loc 2 50 1
	mul.wide.s32 	%rd49, %r1, 4;
	add.s64 	%rd50, %rd48, %rd49;
	ld.global.f32 	%f66, [%rd50];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd50], %f67;
	.loc 2 51 1
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	cvta.to.global.u64 	%rd51, %rd4;
	.loc 2 51 1
	add.s64 	%rd52, %rd51, %rd49;
	ld.global.f32 	%f70, [%rd52];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd52], %f71;
	.loc 2 52 1
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	cvta.to.global.u64 	%rd53, %rd5;
	.loc 2 52 1
	add.s64 	%rd54, %rd53, %rd49;
	ld.global.f32 	%f74, [%rd54];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd54], %f75;

BB0_5:
	.loc 2 54 2
	ret;
}


`
	addslonczewskitorque_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry addslonczewskitorque(
	.param .u64 addslonczewskitorque_param_0,
	.param .u64 addslonczewskitorque_param_1,
	.param .u64 addslonczewskitorque_param_2,
	.param .u64 addslonczewskitorque_param_3,
	.param .u64 addslonczewskitorque_param_4,
	.param .u64 addslonczewskitorque_param_5,
	.param .u64 addslonczewskitorque_param_6,
	.param .u64 addslonczewskitorque_param_7,
	.param .u64 addslonczewskitorque_param_8,
	.param .u64 addslonczewskitorque_param_9,
	.param .u64 addslonczewskitorque_param_10,
	.param .u64 addslonczewskitorque_param_11,
	.param .f32 addslonczewskitorque_param_12,
	.param .u64 addslonczewskitorque_param_13,
	.param .u64 addslonczewskitorque_param_14,
	.param .u64 addslonczewskitorque_param_15,
	.param .u64 addslonczewskitorque_param_16,
	.param .u32 addslonczewskitorque_param_17
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<28>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<55>;
	.reg .f64 	%fd<3>;


	ld.param.u64 	%rd4, [addslonczewskitorque_param_0];
	ld.param.u64 	%rd5, [addslonczewskitorque_param_1];
	ld.param.u64 	%rd6, [addslonczewskitorque_param_2];
	ld.param.u64 	%rd7, [addslonczewskitorque_param_3];
	ld.param.u64 	%rd8, [addslonczewskitorque_param_4];
	ld.param.u64 	%rd9, [addslonczewskitorque_param_5];
	ld.param.u64 	%rd10, [addslonczewskitorque_param_6];
	ld.param.u64 	%rd11, [addslonczewskitorque_param_7];
	ld.param.u64 	%rd12, [addslonczewskitorque_param_8];
	ld.param.u64 	%rd13, [addslonczewskitorque_param_9];
	ld.param.u64 	%rd14, [addslonczewskitorque_param_10];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_11];
	ld.param.f32 	%f19, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd18, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd19, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	cvta.to.global.u64 	%rd1, %rd19;
	.loc 2 16 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 17 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_5;

	cvta.to.global.u64 	%rd20, %rd7;
	.loc 2 19 1
	cvt.s64.s32 	%rd2, %r1;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd20, %rd21;
	ld.global.f32 	%f1, [%rd22];
	cvta.to.global.u64 	%rd23, %rd8;
	.loc 2 19 1
	add.s64 	%rd24, %rd23, %rd21;
	ld.global.f32 	%f2, [%rd24];
	cvta.to.global.u64 	%rd25, %rd9;
	.loc 2 19 1
	add.s64 	%rd26, %rd25, %rd21;
	ld.global.f32 	%f3, [%rd26];
	cvta.to.global.u64 	%rd27, %rd10;
	.loc 2 20 1
	add.s64 	%rd28, %rd27, %rd21;
	ld.global.f32 	%f4, [%rd28];
	.loc 2 23 1
	add.s64 	%rd29, %rd1, %rd2;
	.loc 2 25 1
	ld.global.u8 	%rd3, [%rd29];
	cvta.to.global.u64 	%rd30, %rd11;
	.loc 2 25 1
	shl.b64 	%rd31, %rd3, 2;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd12;
	.loc 2 25 1
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd13;
	.loc 2 25 1
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.f32 	%f5, [%rd32];
	ld.global.f32 	%f6, [%rd34];
	mul.f32 	%f21, %f6, %f6;
	fma.rn.f32 	%f22, %f5, %f5, %f21;
	ld.global.f32 	%f7, [%rd36];
	fma.rn.f32 	%f23, %f7, %f7, %f22;
	.loc 3 991 5
	sqrt.rn.f32 	%f8, %f23;
	mov.f32 	%f76, 0f00000000;
	.loc 2 25 1
	setp.eq.f32 	%p2, %f8, 0f00000000;
	@%p2 bra 	BB0_3;

	rcp.rn.f32 	%f76, %f8;

BB0_3:
	mul.f32 	%f11, %f76, %f5;
	mul.f32 	%f12, %f76, %f6;
	mul.f32 	%f13, %f76, %f7;
	cvta.to.global.u64 	%rd37, %rd14;
	.loc 2 26 1
	add.s64 	%rd39, %rd37, %rd31;
	cvta.to.global.u64 	%rd40, %rd15;
	.loc 2 27 1
	add.s64 	%rd41, %rd40, %rd31;
	ld.global.f32 	%f14, [%rd41];
	cvta.to.global.u64 	%rd42, %rd16;
	.loc 2 28 1
	add.s64 	%rd43, %rd42, %rd31;
	ld.global.f32 	%f15, [%rd43];
	cvta.to.global.u64 	%rd44, %rd17;
	.loc 2 29 1
	add.s64 	%rd45, %rd44, %rd31;
	ld.global.f32 	%f16, [%rd45];
	cvta.to.global.u64 	%rd46, %rd18;
	.loc 2 30 1
	add.s64 	%rd47, %rd46, %rd31;
	ld.global.f32 	%f17, [%rd47];
	.loc 2 26 1
	ld.global.f32 	%f18, [%rd39];
	.loc 2 32 1
	setp.eq.f32 	%p3, %f18, 0f00000000;
	setp.eq.f32 	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB0_5;

	.loc 2 36 1
	mul.f32 	%f24, %f18, %f19;
	.loc 4 2399 3
	div.rn.f32 	%f25, %f4, %f24;
	.loc 2 36 1
	cvt.f64.f32 	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64 	%f26, %fd2;
	.loc 2 37 1
	mul.f32 	%f27, %f16, %f16;
	.loc 2 38 1
	mul.f32 	%f28, %f15, %f27;
	fma.rn.f32 	%f29, %f16, %f16, 0f3F800000;
	fma.rn.f32 	%f30, %f16, %f16, 0fBF800000;
	mul.f32 	%f31, %f12, %f2;
	fma.rn.f32 	%f32, %f11, %f1, %f31;
	fma.rn.f32 	%f33, %f13, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	.loc 4 2399 3
	div.rn.f32 	%f35, %f28, %f34;
	.loc 2 40 1
	mul.f32 	%f36, %f26, %f35;
	.loc 2 41 1
	mul.f32 	%f37, %f26, %f17;
	.loc 2 43 1
	fma.rn.f32 	%f38, %f14, %f14, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	.loc 2 44 1
	mul.f32 	%f40, %f14, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	.loc 2 45 1
	mul.f32 	%f43, %f14, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	.loc 2 47 1
	mul.f32 	%f46, %f13, %f2;
	mul.f32 	%f47, %f12, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f11, %f3;
	mul.f32 	%f50, %f13, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f12, %f1;
	mul.f32 	%f53, %f11, %f2;
	sub.f32 	%f54, %f53, %f52;
	.loc 2 48 1
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	.loc 2 50 1
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	cvta.to.global.u64 	%rd48, %rd4;
	.loc 2 50 1
	shl.b64 	%rd49, %rd2, 2;
	add.s64 	%rd50, %rd48, %rd49;
	ld.global.f32 	%f66, [%rd50];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd50], %f67;
	.loc 2 51 1
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	cvta.to.global.u64 	%rd51, %rd5;
	.loc 2 51 1
	add.s64 	%rd52, %rd51, %rd49;
	ld.global.f32 	%f70, [%rd52];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd52], %f71;
	.loc 2 52 1
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	cvta.to.global.u64 	%rd53, %rd6;
	.loc 2 52 1
	add.s64 	%rd54, %rd53, %rd49;
	ld.global.f32 	%f74, [%rd54];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd54], %f75;

BB0_5:
	.loc 2 54 2
	ret;
}


`
	addslonczewskitorque_ptx_35 = `
.version 3.1
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

.visible .entry addslonczewskitorque(
	.param .u64 addslonczewskitorque_param_0,
	.param .u64 addslonczewskitorque_param_1,
	.param .u64 addslonczewskitorque_param_2,
	.param .u64 addslonczewskitorque_param_3,
	.param .u64 addslonczewskitorque_param_4,
	.param .u64 addslonczewskitorque_param_5,
	.param .u64 addslonczewskitorque_param_6,
	.param .u64 addslonczewskitorque_param_7,
	.param .u64 addslonczewskitorque_param_8,
	.param .u64 addslonczewskitorque_param_9,
	.param .u64 addslonczewskitorque_param_10,
	.param .u64 addslonczewskitorque_param_11,
	.param .f32 addslonczewskitorque_param_12,
	.param .u64 addslonczewskitorque_param_13,
	.param .u64 addslonczewskitorque_param_14,
	.param .u64 addslonczewskitorque_param_15,
	.param .u64 addslonczewskitorque_param_16,
	.param .u32 addslonczewskitorque_param_17
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<13>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<55>;
	.reg .f64 	%fd<3>;


	ld.param.u64 	%rd4, [addslonczewskitorque_param_0];
	ld.param.u64 	%rd5, [addslonczewskitorque_param_1];
	ld.param.u64 	%rd6, [addslonczewskitorque_param_2];
	ld.param.u64 	%rd7, [addslonczewskitorque_param_3];
	ld.param.u64 	%rd8, [addslonczewskitorque_param_4];
	ld.param.u64 	%rd9, [addslonczewskitorque_param_5];
	ld.param.u64 	%rd10, [addslonczewskitorque_param_6];
	ld.param.u64 	%rd11, [addslonczewskitorque_param_7];
	ld.param.u64 	%rd12, [addslonczewskitorque_param_8];
	ld.param.u64 	%rd13, [addslonczewskitorque_param_9];
	ld.param.u64 	%rd14, [addslonczewskitorque_param_10];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_11];
	ld.param.f32 	%f19, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd18, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd19, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	cvta.to.global.u64 	%rd1, %rd19;
	.loc 3 16 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 3 17 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB2_5;

	cvta.to.global.u64 	%rd20, %rd7;
	.loc 3 19 1
	cvt.s64.s32 	%rd2, %r1;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd20, %rd21;
	ld.global.nc.f32 	%f1, [%rd22];
	cvta.to.global.u64 	%rd23, %rd8;
	.loc 3 19 1
	add.s64 	%rd24, %rd23, %rd21;
	ld.global.nc.f32 	%f2, [%rd24];
	cvta.to.global.u64 	%rd25, %rd9;
	.loc 3 19 1
	add.s64 	%rd26, %rd25, %rd21;
	ld.global.nc.f32 	%f3, [%rd26];
	cvta.to.global.u64 	%rd27, %rd10;
	.loc 3 20 1
	add.s64 	%rd28, %rd27, %rd21;
	ld.global.nc.f32 	%f4, [%rd28];
	.loc 3 23 1
	add.s64 	%rd29, %rd1, %rd2;
	.loc 3 25 1
	ld.global.u8 	%rd3, [%rd29];
	cvta.to.global.u64 	%rd30, %rd11;
	.loc 3 25 1
	shl.b64 	%rd31, %rd3, 2;
	add.s64 	%rd32, %rd30, %rd31;
	ld.global.nc.f32 	%f5, [%rd32];
	cvta.to.global.u64 	%rd33, %rd12;
	.loc 3 25 1
	add.s64 	%rd34, %rd33, %rd31;
	ld.global.nc.f32 	%f6, [%rd34];
	cvta.to.global.u64 	%rd35, %rd13;
	.loc 3 25 1
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.nc.f32 	%f7, [%rd36];
	mul.f32 	%f21, %f6, %f6;
	fma.rn.f32 	%f22, %f5, %f5, %f21;
	fma.rn.f32 	%f23, %f7, %f7, %f22;
	.loc 4 991 5
	sqrt.rn.f32 	%f8, %f23;
	mov.f32 	%f76, 0f00000000;
	.loc 3 25 1
	setp.eq.f32 	%p2, %f8, 0f00000000;
	@%p2 bra 	BB2_3;

	rcp.rn.f32 	%f76, %f8;

BB2_3:
	mul.f32 	%f11, %f76, %f5;
	mul.f32 	%f12, %f76, %f6;
	mul.f32 	%f13, %f76, %f7;
	cvta.to.global.u64 	%rd37, %rd14;
	.loc 3 26 1
	add.s64 	%rd39, %rd37, %rd31;
	ld.global.nc.f32 	%f14, [%rd39];
	cvta.to.global.u64 	%rd40, %rd15;
	.loc 3 27 1
	add.s64 	%rd41, %rd40, %rd31;
	ld.global.nc.f32 	%f15, [%rd41];
	cvta.to.global.u64 	%rd42, %rd16;
	.loc 3 28 1
	add.s64 	%rd43, %rd42, %rd31;
	ld.global.nc.f32 	%f16, [%rd43];
	cvta.to.global.u64 	%rd44, %rd17;
	.loc 3 29 1
	add.s64 	%rd45, %rd44, %rd31;
	ld.global.nc.f32 	%f17, [%rd45];
	cvta.to.global.u64 	%rd46, %rd18;
	.loc 3 30 1
	add.s64 	%rd47, %rd46, %rd31;
	ld.global.nc.f32 	%f18, [%rd47];
	.loc 3 32 1
	setp.eq.f32 	%p3, %f14, 0f00000000;
	setp.eq.f32 	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB2_5;

	.loc 3 36 1
	mul.f32 	%f24, %f14, %f19;
	.loc 5 2399 3
	div.rn.f32 	%f25, %f4, %f24;
	.loc 3 36 1
	cvt.f64.f32 	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64 	%f26, %fd2;
	.loc 3 37 1
	mul.f32 	%f27, %f17, %f17;
	.loc 3 38 1
	mul.f32 	%f28, %f16, %f27;
	fma.rn.f32 	%f29, %f17, %f17, 0f3F800000;
	fma.rn.f32 	%f30, %f17, %f17, 0fBF800000;
	mul.f32 	%f31, %f12, %f2;
	fma.rn.f32 	%f32, %f11, %f1, %f31;
	fma.rn.f32 	%f33, %f13, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	.loc 5 2399 3
	div.rn.f32 	%f35, %f28, %f34;
	.loc 3 40 1
	mul.f32 	%f36, %f26, %f35;
	.loc 3 41 1
	mul.f32 	%f37, %f26, %f18;
	.loc 3 43 1
	fma.rn.f32 	%f38, %f15, %f15, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	.loc 3 44 1
	mul.f32 	%f40, %f15, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	.loc 3 45 1
	mul.f32 	%f43, %f15, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	.loc 3 47 1
	mul.f32 	%f46, %f13, %f2;
	mul.f32 	%f47, %f12, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f11, %f3;
	mul.f32 	%f50, %f13, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f12, %f1;
	mul.f32 	%f53, %f11, %f2;
	sub.f32 	%f54, %f53, %f52;
	.loc 3 48 1
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	.loc 3 50 1
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	cvta.to.global.u64 	%rd48, %rd4;
	.loc 3 50 1
	shl.b64 	%rd49, %rd2, 2;
	add.s64 	%rd50, %rd48, %rd49;
	ld.global.nc.f32 	%f66, [%rd50];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd50], %f67;
	.loc 3 51 1
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	cvta.to.global.u64 	%rd51, %rd5;
	.loc 3 51 1
	add.s64 	%rd52, %rd51, %rd49;
	ld.global.nc.f32 	%f70, [%rd52];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd52], %f71;
	.loc 3 52 1
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	cvta.to.global.u64 	%rd53, %rd6;
	.loc 3 52 1
	add.s64 	%rd54, %rd53, %rd49;
	ld.global.nc.f32 	%f74, [%rd54];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd54], %f75;

BB2_5:
	.loc 3 54 2
	ret;
}


`
)
