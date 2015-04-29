package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
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
		timer.Start("addslonczewskitorque")
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
		timer.Stop("addslonczewskitorque")
	}
}

// maps compute capability on PTX code for addslonczewskitorque kernel.
var addslonczewskitorque_map = map[int]string{0: "",
	20: addslonczewskitorque_ptx_20,
	30: addslonczewskitorque_ptx_30,
	35: addslonczewskitorque_ptx_35,
	50: addslonczewskitorque_ptx_50}

// addslonczewskitorque PTX code for various compute capabilities.
const (
	addslonczewskitorque_ptx_20 = `
.version 4.0
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
	.reg .s32 	%r<16>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<56>;
	.reg .f64 	%fd<3>;


	ld.param.u64 	%rd2, [addslonczewskitorque_param_0];
	ld.param.u64 	%rd3, [addslonczewskitorque_param_1];
	ld.param.u64 	%rd4, [addslonczewskitorque_param_2];
	ld.param.u64 	%rd5, [addslonczewskitorque_param_3];
	ld.param.u64 	%rd6, [addslonczewskitorque_param_4];
	ld.param.u64 	%rd7, [addslonczewskitorque_param_5];
	ld.param.u64 	%rd8, [addslonczewskitorque_param_6];
	ld.param.u64 	%rd9, [addslonczewskitorque_param_7];
	ld.param.u64 	%rd10, [addslonczewskitorque_param_8];
	ld.param.u64 	%rd11, [addslonczewskitorque_param_9];
	ld.param.u64 	%rd12, [addslonczewskitorque_param_10];
	ld.param.u64 	%rd13, [addslonczewskitorque_param_11];
	ld.param.f32 	%f15, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd14, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_6;

	cvta.to.global.u64 	%rd18, %rd17;
	cvta.to.global.u64 	%rd19, %rd5;
	cvt.s64.s32	%rd20, %r1;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd19, %rd21;
	ld.global.f32 	%f1, [%rd22];
	cvta.to.global.u64 	%rd23, %rd6;
	add.s64 	%rd24, %rd23, %rd21;
	ld.global.f32 	%f2, [%rd24];
	cvta.to.global.u64 	%rd25, %rd7;
	add.s64 	%rd26, %rd25, %rd21;
	ld.global.f32 	%f3, [%rd26];
	cvta.to.global.u64 	%rd27, %rd8;
	add.s64 	%rd28, %rd27, %rd21;
	ld.global.f32 	%f4, [%rd28];
	add.s64 	%rd29, %rd18, %rd20;
	ld.global.u8 	%rd1, [%rd29];
	cvta.to.global.u64 	%rd30, %rd9;
	shl.b64 	%rd31, %rd1, 2;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd10;
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd11;
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.f32 	%f5, [%rd32];
	ld.global.f32 	%f6, [%rd34];
	mul.f32 	%f16, %f6, %f6;
	fma.rn.f32 	%f17, %f5, %f5, %f16;
	ld.global.f32 	%f7, [%rd36];
	fma.rn.f32 	%f18, %f7, %f7, %f17;
	sqrt.rn.f32 	%f8, %f18;
	setp.neu.f32	%p2, %f8, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f76, %f8;

BB0_4:
	mul.f32 	%f11, %f76, %f5;
	mul.f32 	%f12, %f76, %f6;
	mul.f32 	%f13, %f76, %f7;
	cvta.to.global.u64 	%rd37, %rd12;
	add.s64 	%rd39, %rd37, %rd31;
	ld.global.f32 	%f14, [%rd39];
	setp.eq.f32	%p3, %f14, 0f00000000;
	setp.eq.f32	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB0_6;

	cvta.to.global.u64 	%rd40, %rd13;
	add.s64 	%rd42, %rd40, %rd31;
	ld.global.f32 	%f20, [%rd42];
	cvta.to.global.u64 	%rd43, %rd14;
	add.s64 	%rd44, %rd43, %rd31;
	ld.global.f32 	%f21, [%rd44];
	cvta.to.global.u64 	%rd45, %rd15;
	add.s64 	%rd46, %rd45, %rd31;
	ld.global.f32 	%f22, [%rd46];
	cvta.to.global.u64 	%rd47, %rd16;
	add.s64 	%rd48, %rd47, %rd31;
	ld.global.f32 	%f23, [%rd48];
	mul.f32 	%f24, %f14, %f15;
	div.rn.f32 	%f25, %f4, %f24;
	cvt.f64.f32	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64	%f26, %fd2;
	mul.f32 	%f27, %f22, %f22;
	mul.f32 	%f28, %f21, %f27;
	add.f32 	%f29, %f27, 0f3F800000;
	add.f32 	%f30, %f27, 0fBF800000;
	mul.f32 	%f31, %f12, %f2;
	fma.rn.f32 	%f32, %f11, %f1, %f31;
	fma.rn.f32 	%f33, %f13, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	div.rn.f32 	%f35, %f28, %f34;
	mul.f32 	%f36, %f26, %f35;
	mul.f32 	%f37, %f26, %f23;
	fma.rn.f32 	%f38, %f20, %f20, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	mul.f32 	%f40, %f20, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	mul.f32 	%f43, %f20, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	mul.f32 	%f46, %f13, %f2;
	mul.f32 	%f47, %f12, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f11, %f3;
	mul.f32 	%f50, %f13, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f12, %f1;
	mul.f32 	%f53, %f11, %f2;
	sub.f32 	%f54, %f53, %f52;
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	cvta.to.global.u64 	%rd49, %rd2;
	mul.wide.s32 	%rd50, %r1, 4;
	add.s64 	%rd51, %rd49, %rd50;
	ld.global.f32 	%f66, [%rd51];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd51], %f67;
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	cvta.to.global.u64 	%rd52, %rd3;
	add.s64 	%rd53, %rd52, %rd50;
	ld.global.f32 	%f70, [%rd53];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd53], %f71;
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	cvta.to.global.u64 	%rd54, %rd4;
	add.s64 	%rd55, %rd54, %rd50;
	ld.global.f32 	%f74, [%rd55];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd55], %f75;

BB0_6:
	ret;
}


`
	addslonczewskitorque_ptx_30 = `
.version 4.0
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
	.reg .s32 	%r<16>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<56>;
	.reg .f64 	%fd<3>;


	ld.param.u64 	%rd2, [addslonczewskitorque_param_0];
	ld.param.u64 	%rd3, [addslonczewskitorque_param_1];
	ld.param.u64 	%rd4, [addslonczewskitorque_param_2];
	ld.param.u64 	%rd5, [addslonczewskitorque_param_3];
	ld.param.u64 	%rd6, [addslonczewskitorque_param_4];
	ld.param.u64 	%rd7, [addslonczewskitorque_param_5];
	ld.param.u64 	%rd8, [addslonczewskitorque_param_6];
	ld.param.u64 	%rd9, [addslonczewskitorque_param_7];
	ld.param.u64 	%rd10, [addslonczewskitorque_param_8];
	ld.param.u64 	%rd11, [addslonczewskitorque_param_9];
	ld.param.u64 	%rd12, [addslonczewskitorque_param_10];
	ld.param.u64 	%rd13, [addslonczewskitorque_param_11];
	ld.param.f32 	%f15, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd14, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_6;

	cvta.to.global.u64 	%rd18, %rd17;
	cvta.to.global.u64 	%rd19, %rd5;
	cvt.s64.s32	%rd20, %r1;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd19, %rd21;
	ld.global.f32 	%f1, [%rd22];
	cvta.to.global.u64 	%rd23, %rd6;
	add.s64 	%rd24, %rd23, %rd21;
	ld.global.f32 	%f2, [%rd24];
	cvta.to.global.u64 	%rd25, %rd7;
	add.s64 	%rd26, %rd25, %rd21;
	ld.global.f32 	%f3, [%rd26];
	cvta.to.global.u64 	%rd27, %rd8;
	add.s64 	%rd28, %rd27, %rd21;
	ld.global.f32 	%f4, [%rd28];
	add.s64 	%rd29, %rd18, %rd20;
	ld.global.u8 	%rd1, [%rd29];
	cvta.to.global.u64 	%rd30, %rd9;
	shl.b64 	%rd31, %rd1, 2;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd10;
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd11;
	add.s64 	%rd36, %rd35, %rd31;
	ld.global.f32 	%f5, [%rd32];
	ld.global.f32 	%f6, [%rd34];
	mul.f32 	%f16, %f6, %f6;
	fma.rn.f32 	%f17, %f5, %f5, %f16;
	ld.global.f32 	%f7, [%rd36];
	fma.rn.f32 	%f18, %f7, %f7, %f17;
	sqrt.rn.f32 	%f8, %f18;
	setp.neu.f32	%p2, %f8, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f76, %f8;

BB0_4:
	mul.f32 	%f11, %f76, %f5;
	mul.f32 	%f12, %f76, %f6;
	mul.f32 	%f13, %f76, %f7;
	cvta.to.global.u64 	%rd37, %rd12;
	add.s64 	%rd39, %rd37, %rd31;
	ld.global.f32 	%f14, [%rd39];
	setp.eq.f32	%p3, %f14, 0f00000000;
	setp.eq.f32	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB0_6;

	cvta.to.global.u64 	%rd40, %rd13;
	add.s64 	%rd42, %rd40, %rd31;
	ld.global.f32 	%f20, [%rd42];
	cvta.to.global.u64 	%rd43, %rd14;
	add.s64 	%rd44, %rd43, %rd31;
	ld.global.f32 	%f21, [%rd44];
	cvta.to.global.u64 	%rd45, %rd15;
	add.s64 	%rd46, %rd45, %rd31;
	ld.global.f32 	%f22, [%rd46];
	cvta.to.global.u64 	%rd47, %rd16;
	add.s64 	%rd48, %rd47, %rd31;
	ld.global.f32 	%f23, [%rd48];
	mul.f32 	%f24, %f14, %f15;
	div.rn.f32 	%f25, %f4, %f24;
	cvt.f64.f32	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64	%f26, %fd2;
	mul.f32 	%f27, %f22, %f22;
	mul.f32 	%f28, %f21, %f27;
	add.f32 	%f29, %f27, 0f3F800000;
	add.f32 	%f30, %f27, 0fBF800000;
	mul.f32 	%f31, %f12, %f2;
	fma.rn.f32 	%f32, %f11, %f1, %f31;
	fma.rn.f32 	%f33, %f13, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	div.rn.f32 	%f35, %f28, %f34;
	mul.f32 	%f36, %f26, %f35;
	mul.f32 	%f37, %f26, %f23;
	fma.rn.f32 	%f38, %f20, %f20, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	mul.f32 	%f40, %f20, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	mul.f32 	%f43, %f20, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	mul.f32 	%f46, %f13, %f2;
	mul.f32 	%f47, %f12, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f11, %f3;
	mul.f32 	%f50, %f13, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f12, %f1;
	mul.f32 	%f53, %f11, %f2;
	sub.f32 	%f54, %f53, %f52;
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	cvta.to.global.u64 	%rd49, %rd2;
	mul.wide.s32 	%rd50, %r1, 4;
	add.s64 	%rd51, %rd49, %rd50;
	ld.global.f32 	%f66, [%rd51];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd51], %f67;
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	cvta.to.global.u64 	%rd52, %rd3;
	add.s64 	%rd53, %rd52, %rd50;
	ld.global.f32 	%f70, [%rd53];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd53], %f71;
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	cvta.to.global.u64 	%rd54, %rd4;
	add.s64 	%rd55, %rd54, %rd50;
	ld.global.f32 	%f74, [%rd55];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd55], %f75;

BB0_6:
	ret;
}


`
	addslonczewskitorque_ptx_35 = `
.version 4.1
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
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
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
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<77>;
	.reg .s64 	%rd<57>;
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
	ld.param.f32 	%f12, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd18, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_6;

	cvta.to.global.u64 	%rd19, %rd12;
	cvta.to.global.u64 	%rd20, %rd11;
	cvta.to.global.u64 	%rd21, %rd10;
	cvta.to.global.u64 	%rd22, %rd18;
	cvta.to.global.u64 	%rd23, %rd9;
	cvta.to.global.u64 	%rd24, %rd8;
	cvta.to.global.u64 	%rd25, %rd7;
	cvta.to.global.u64 	%rd26, %rd6;
	cvt.s64.s32	%rd1, %r1;
	mul.wide.s32 	%rd27, %r1, 4;
	add.s64 	%rd28, %rd26, %rd27;
	ld.global.nc.f32 	%f1, [%rd28];
	add.s64 	%rd29, %rd25, %rd27;
	ld.global.nc.f32 	%f2, [%rd29];
	add.s64 	%rd30, %rd24, %rd27;
	ld.global.nc.f32 	%f3, [%rd30];
	add.s64 	%rd31, %rd23, %rd27;
	ld.global.nc.f32 	%f4, [%rd31];
	add.s64 	%rd32, %rd22, %rd1;
	ld.global.nc.u8 	%rs1, [%rd32];
	cvt.u64.u16	%rd33, %rs1;
	and.b64  	%rd2, %rd33, 255;
	shl.b64 	%rd34, %rd2, 2;
	add.s64 	%rd35, %rd21, %rd34;
	add.s64 	%rd36, %rd20, %rd34;
	add.s64 	%rd37, %rd19, %rd34;
	ld.global.nc.f32 	%f5, [%rd35];
	ld.global.nc.f32 	%f6, [%rd36];
	mul.f32 	%f13, %f6, %f6;
	fma.rn.f32 	%f14, %f5, %f5, %f13;
	ld.global.nc.f32 	%f7, [%rd37];
	fma.rn.f32 	%f15, %f7, %f7, %f14;
	sqrt.rn.f32 	%f8, %f15;
	setp.neu.f32	%p2, %f8, 0f00000000;
	@%p2 bra 	BB5_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB5_4;

BB5_3:
	rcp.rn.f32 	%f76, %f8;

BB5_4:
	cvta.to.global.u64 	%rd38, %rd13;
	add.s64 	%rd40, %rd38, %rd34;
	ld.global.nc.f32 	%f11, [%rd40];
	setp.eq.f32	%p3, %f11, 0f00000000;
	setp.eq.f32	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB5_6;

	cvta.to.global.u64 	%rd41, %rd5;
	cvta.to.global.u64 	%rd42, %rd4;
	cvta.to.global.u64 	%rd43, %rd3;
	cvta.to.global.u64 	%rd44, %rd17;
	cvta.to.global.u64 	%rd45, %rd16;
	cvta.to.global.u64 	%rd46, %rd15;
	mul.f32 	%f17, %f76, %f5;
	mul.f32 	%f18, %f76, %f6;
	mul.f32 	%f19, %f76, %f7;
	cvta.to.global.u64 	%rd47, %rd14;
	add.s64 	%rd49, %rd47, %rd34;
	ld.global.nc.f32 	%f20, [%rd49];
	add.s64 	%rd50, %rd46, %rd34;
	ld.global.nc.f32 	%f21, [%rd50];
	add.s64 	%rd51, %rd45, %rd34;
	ld.global.nc.f32 	%f22, [%rd51];
	add.s64 	%rd52, %rd44, %rd34;
	ld.global.nc.f32 	%f23, [%rd52];
	mul.f32 	%f24, %f11, %f12;
	div.rn.f32 	%f25, %f4, %f24;
	cvt.f64.f32	%fd1, %f25;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64	%f26, %fd2;
	mul.f32 	%f27, %f22, %f22;
	mul.f32 	%f28, %f21, %f27;
	add.f32 	%f29, %f27, 0f3F800000;
	add.f32 	%f30, %f27, 0fBF800000;
	mul.f32 	%f31, %f18, %f2;
	fma.rn.f32 	%f32, %f17, %f1, %f31;
	fma.rn.f32 	%f33, %f19, %f3, %f32;
	fma.rn.f32 	%f34, %f30, %f33, %f29;
	div.rn.f32 	%f35, %f28, %f34;
	mul.f32 	%f36, %f26, %f35;
	mul.f32 	%f37, %f26, %f23;
	fma.rn.f32 	%f38, %f20, %f20, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	mul.f32 	%f40, %f20, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	mul.f32 	%f43, %f20, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	mul.f32 	%f46, %f19, %f2;
	mul.f32 	%f47, %f18, %f3;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f17, %f3;
	mul.f32 	%f50, %f19, %f1;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f18, %f1;
	mul.f32 	%f53, %f17, %f2;
	sub.f32 	%f54, %f53, %f52;
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	mul.f32 	%f64, %f45, %f48;
	fma.rn.f32 	%f65, %f42, %f57, %f64;
	shl.b64 	%rd53, %rd1, 2;
	add.s64 	%rd54, %rd43, %rd53;
	ld.global.f32 	%f66, [%rd54];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd54], %f67;
	mul.f32 	%f68, %f45, %f51;
	fma.rn.f32 	%f69, %f42, %f60, %f68;
	add.s64 	%rd55, %rd42, %rd53;
	ld.global.f32 	%f70, [%rd55];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd55], %f71;
	mul.f32 	%f72, %f45, %f54;
	fma.rn.f32 	%f73, %f42, %f63, %f72;
	add.s64 	%rd56, %rd41, %rd53;
	ld.global.f32 	%f74, [%rd56];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd56], %f75;

BB5_6:
	ret;
}


`
	addslonczewskitorque_ptx_50 = `
.version 4.2
.target sm_50
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	addslonczewskitorque
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
	.reg .s16 	%rs<2>;
	.reg .f32 	%f<77>;
	.reg .s32 	%r<11>;
	.reg .f64 	%fd<3>;
	.reg .s64 	%rd<57>;


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
	ld.param.f32 	%f12, [addslonczewskitorque_param_12];
	ld.param.u64 	%rd15, [addslonczewskitorque_param_13];
	ld.param.u64 	%rd16, [addslonczewskitorque_param_14];
	ld.param.u64 	%rd17, [addslonczewskitorque_param_15];
	ld.param.u64 	%rd18, [addslonczewskitorque_param_16];
	ld.param.u32 	%r2, [addslonczewskitorque_param_17];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_5;

	cvta.to.global.u64 	%rd19, %rd6;
	cvt.s64.s32	%rd1, %r1;
	mul.wide.s32 	%rd20, %r1, 4;
	add.s64 	%rd21, %rd19, %rd20;
	ld.global.nc.f32 	%f1, [%rd21];
	cvta.to.global.u64 	%rd22, %rd7;
	add.s64 	%rd23, %rd22, %rd20;
	ld.global.nc.f32 	%f2, [%rd23];
	cvta.to.global.u64 	%rd24, %rd8;
	add.s64 	%rd25, %rd24, %rd20;
	ld.global.nc.f32 	%f3, [%rd25];
	cvta.to.global.u64 	%rd26, %rd9;
	add.s64 	%rd27, %rd26, %rd20;
	ld.global.nc.f32 	%f4, [%rd27];
	cvta.to.global.u64 	%rd28, %rd18;
	add.s64 	%rd29, %rd28, %rd1;
	ld.global.nc.u8 	%rs1, [%rd29];
	cvt.u64.u16	%rd30, %rs1;
	and.b64  	%rd2, %rd30, 255;
	cvta.to.global.u64 	%rd31, %rd10;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd32, %r10, 4;
	add.s64 	%rd33, %rd31, %rd32;
	cvta.to.global.u64 	%rd34, %rd11;
	add.s64 	%rd35, %rd34, %rd32;
	cvta.to.global.u64 	%rd36, %rd12;
	add.s64 	%rd37, %rd36, %rd32;
	ld.global.nc.f32 	%f5, [%rd33];
	ld.global.nc.f32 	%f6, [%rd35];
	mul.f32 	%f14, %f6, %f6;
	fma.rn.f32 	%f15, %f5, %f5, %f14;
	ld.global.nc.f32 	%f7, [%rd37];
	fma.rn.f32 	%f16, %f7, %f7, %f15;
	sqrt.rn.f32 	%f8, %f16;
	mov.f32 	%f76, 0f00000000;
	setp.eq.f32	%p2, %f8, 0f00000000;
	@%p2 bra 	BB6_3;

	rcp.rn.f32 	%f76, %f8;

BB6_3:
	cvta.to.global.u64 	%rd38, %rd13;
	shl.b64 	%rd39, %rd2, 2;
	add.s64 	%rd40, %rd38, %rd39;
	ld.global.nc.f32 	%f11, [%rd40];
	setp.eq.f32	%p3, %f11, 0f00000000;
	setp.eq.f32	%p4, %f4, 0f00000000;
	or.pred  	%p5, %p4, %p3;
	@%p5 bra 	BB6_5;

	cvta.to.global.u64 	%rd41, %rd5;
	cvta.to.global.u64 	%rd42, %rd4;
	cvta.to.global.u64 	%rd43, %rd3;
	cvta.to.global.u64 	%rd44, %rd17;
	cvta.to.global.u64 	%rd45, %rd16;
	cvta.to.global.u64 	%rd46, %rd15;
	mul.f32 	%f17, %f5, %f76;
	mul.f32 	%f18, %f6, %f76;
	mul.f32 	%f19, %f7, %f76;
	cvta.to.global.u64 	%rd47, %rd14;
	add.s64 	%rd49, %rd47, %rd39;
	ld.global.nc.f32 	%f20, [%rd49];
	add.s64 	%rd50, %rd46, %rd39;
	add.s64 	%rd51, %rd45, %rd39;
	ld.global.nc.f32 	%f21, [%rd51];
	add.s64 	%rd52, %rd44, %rd39;
	ld.global.nc.f32 	%f22, [%rd52];
	mul.f32 	%f23, %f11, %f12;
	div.rn.f32 	%f24, %f4, %f23;
	cvt.f64.f32	%fd1, %f24;
	mul.f64 	%fd2, %fd1, 0d3CC7B6EF14E9250C;
	cvt.rn.f32.f64	%f25, %fd2;
	mul.f32 	%f26, %f21, %f21;
	ld.global.nc.f32 	%f27, [%rd50];
	mul.f32 	%f28, %f26, %f27;
	add.f32 	%f29, %f26, 0f3F800000;
	add.f32 	%f30, %f26, 0fBF800000;
	mul.f32 	%f31, %f2, %f18;
	fma.rn.f32 	%f32, %f1, %f17, %f31;
	fma.rn.f32 	%f33, %f3, %f19, %f32;
	fma.rn.f32 	%f34, %f33, %f30, %f29;
	div.rn.f32 	%f35, %f28, %f34;
	mul.f32 	%f36, %f35, %f25;
	mul.f32 	%f37, %f22, %f25;
	fma.rn.f32 	%f38, %f20, %f20, 0f3F800000;
	rcp.rn.f32 	%f39, %f38;
	mul.f32 	%f40, %f20, %f37;
	sub.f32 	%f41, %f36, %f40;
	mul.f32 	%f42, %f39, %f41;
	mul.f32 	%f43, %f20, %f36;
	sub.f32 	%f44, %f37, %f43;
	mul.f32 	%f45, %f39, %f44;
	mul.f32 	%f46, %f2, %f19;
	mul.f32 	%f47, %f3, %f18;
	sub.f32 	%f48, %f47, %f46;
	mul.f32 	%f49, %f3, %f17;
	mul.f32 	%f50, %f1, %f19;
	sub.f32 	%f51, %f50, %f49;
	mul.f32 	%f52, %f1, %f18;
	mul.f32 	%f53, %f2, %f17;
	sub.f32 	%f54, %f53, %f52;
	mul.f32 	%f55, %f2, %f54;
	mul.f32 	%f56, %f3, %f51;
	sub.f32 	%f57, %f55, %f56;
	mul.f32 	%f58, %f3, %f48;
	mul.f32 	%f59, %f1, %f54;
	sub.f32 	%f60, %f58, %f59;
	mul.f32 	%f61, %f1, %f51;
	mul.f32 	%f62, %f2, %f48;
	sub.f32 	%f63, %f61, %f62;
	mul.f32 	%f64, %f48, %f45;
	fma.rn.f32 	%f65, %f57, %f42, %f64;
	shl.b64 	%rd53, %rd1, 2;
	add.s64 	%rd54, %rd43, %rd53;
	ld.global.f32 	%f66, [%rd54];
	add.f32 	%f67, %f66, %f65;
	st.global.f32 	[%rd54], %f67;
	mul.f32 	%f68, %f51, %f45;
	fma.rn.f32 	%f69, %f60, %f42, %f68;
	add.s64 	%rd55, %rd42, %rd53;
	ld.global.f32 	%f70, [%rd55];
	add.f32 	%f71, %f70, %f69;
	st.global.f32 	[%rd55], %f71;
	mul.f32 	%f72, %f54, %f45;
	fma.rn.f32 	%f73, %f63, %f42, %f72;
	add.s64 	%rd56, %rd41, %rd53;
	ld.global.f32 	%f74, [%rd56];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd56], %f75;

BB6_5:
	ret;
}


`
)
