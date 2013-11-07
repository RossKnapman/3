package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var regionaddv_code cu.Function

type regionaddv_args struct {
	arg_dstx    unsafe.Pointer
	arg_dsty    unsafe.Pointer
	arg_dstz    unsafe.Pointer
	arg_LUTx    unsafe.Pointer
	arg_LUTy    unsafe.Pointer
	arg_LUTz    unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [8]unsafe.Pointer
}

// Wrapper for regionaddv CUDA kernel, asynchronous.
func k_regionaddv_async(dstx unsafe.Pointer, dsty unsafe.Pointer, dstz unsafe.Pointer, LUTx unsafe.Pointer, LUTy unsafe.Pointer, LUTz unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config, str int) {
	if synchronous { // debug
		SyncAll()
	}

	if regionaddv_code == 0 {
		regionaddv_code = fatbinLoad(regionaddv_map, "regionaddv")
	}

	var _a_ regionaddv_args

	_a_.arg_dstx = dstx
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dstx)
	_a_.arg_dsty = dsty
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_dsty)
	_a_.arg_dstz = dstz
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_dstz)
	_a_.arg_LUTx = LUTx
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_LUTx)
	_a_.arg_LUTy = LUTy
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_LUTy)
	_a_.arg_LUTz = LUTz
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_LUTz)
	_a_.arg_regions = regions
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_regions)
	_a_.arg_N = N
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_N)

	args := _a_.argptr[:]
	cu.LaunchKernel(regionaddv_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)

	if synchronous { // debug
		SyncAll()
	}
}

// Wrapper for regionaddv CUDA kernel, synchronized.
func k_regionaddv(dstx unsafe.Pointer, dsty unsafe.Pointer, dstz unsafe.Pointer, LUTx unsafe.Pointer, LUTy unsafe.Pointer, LUTz unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	const stream = 0
	k_regionaddv_async(dstx, dsty, dstz, LUTx, LUTy, LUTz, regions, N, cfg, stream)
	Sync(stream)
}

var regionaddv_map = map[int]string{0: "",
	20: regionaddv_ptx_20,
	30: regionaddv_ptx_30,
	35: regionaddv_ptx_35}

const (
	regionaddv_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry regionaddv(
	.param .u64 regionaddv_param_0,
	.param .u64 regionaddv_param_1,
	.param .u64 regionaddv_param_2,
	.param .u64 regionaddv_param_3,
	.param .u64 regionaddv_param_4,
	.param .u64 regionaddv_param_5,
	.param .u64 regionaddv_param_6,
	.param .u32 regionaddv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<10>;
	.reg .s64 	%rd<26>;


	ld.param.u64 	%rd8, [regionaddv_param_0];
	ld.param.u64 	%rd9, [regionaddv_param_1];
	ld.param.u64 	%rd10, [regionaddv_param_2];
	ld.param.u64 	%rd11, [regionaddv_param_3];
	ld.param.u64 	%rd12, [regionaddv_param_4];
	ld.param.u64 	%rd13, [regionaddv_param_5];
	ld.param.u64 	%rd14, [regionaddv_param_6];
	ld.param.u32 	%r2, [regionaddv_param_7];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd13;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd11;
	cvta.to.global.u64 	%rd7, %rd14;
	.loc 1 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 10 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvt.s64.s32	%rd15, %r1;
	add.s64 	%rd16, %rd7, %rd15;
	ld.global.s8 	%rd17, [%rd16];
	shl.b64 	%rd18, %rd17, 2;
	add.s64 	%rd19, %rd6, %rd18;
	mul.wide.s32 	%rd20, %r1, 4;
	add.s64 	%rd21, %rd5, %rd20;
	.loc 1 13 1
	ld.global.f32 	%f1, [%rd21];
	ld.global.f32 	%f2, [%rd19];
	add.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd21], %f3;
	add.s64 	%rd22, %rd4, %rd18;
	add.s64 	%rd23, %rd3, %rd20;
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd23];
	ld.global.f32 	%f5, [%rd22];
	add.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd23], %f6;
	add.s64 	%rd24, %rd2, %rd18;
	add.s64 	%rd25, %rd1, %rd20;
	.loc 1 15 1
	ld.global.f32 	%f7, [%rd25];
	ld.global.f32 	%f8, [%rd24];
	add.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd25], %f9;

BB0_2:
	.loc 1 17 2
	ret;
}


`
	regionaddv_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry regionaddv(
	.param .u64 regionaddv_param_0,
	.param .u64 regionaddv_param_1,
	.param .u64 regionaddv_param_2,
	.param .u64 regionaddv_param_3,
	.param .u64 regionaddv_param_4,
	.param .u64 regionaddv_param_5,
	.param .u64 regionaddv_param_6,
	.param .u32 regionaddv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<10>;
	.reg .s64 	%rd<26>;


	ld.param.u64 	%rd8, [regionaddv_param_0];
	ld.param.u64 	%rd9, [regionaddv_param_1];
	ld.param.u64 	%rd10, [regionaddv_param_2];
	ld.param.u64 	%rd11, [regionaddv_param_3];
	ld.param.u64 	%rd12, [regionaddv_param_4];
	ld.param.u64 	%rd13, [regionaddv_param_5];
	ld.param.u64 	%rd14, [regionaddv_param_6];
	ld.param.u32 	%r2, [regionaddv_param_7];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd13;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd11;
	cvta.to.global.u64 	%rd7, %rd14;
	.loc 1 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 10 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvt.s64.s32	%rd15, %r1;
	add.s64 	%rd16, %rd7, %rd15;
	ld.global.s8 	%rd17, [%rd16];
	shl.b64 	%rd18, %rd17, 2;
	add.s64 	%rd19, %rd6, %rd18;
	mul.wide.s32 	%rd20, %r1, 4;
	add.s64 	%rd21, %rd5, %rd20;
	.loc 1 13 1
	ld.global.f32 	%f1, [%rd21];
	ld.global.f32 	%f2, [%rd19];
	add.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd21], %f3;
	add.s64 	%rd22, %rd4, %rd18;
	add.s64 	%rd23, %rd3, %rd20;
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd23];
	ld.global.f32 	%f5, [%rd22];
	add.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd23], %f6;
	add.s64 	%rd24, %rd2, %rd18;
	add.s64 	%rd25, %rd1, %rd20;
	.loc 1 15 1
	ld.global.f32 	%f7, [%rd25];
	ld.global.f32 	%f8, [%rd24];
	add.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd25], %f9;

BB0_2:
	.loc 1 17 2
	ret;
}


`
	regionaddv_ptx_35 = `
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

.visible .entry regionaddv(
	.param .u64 regionaddv_param_0,
	.param .u64 regionaddv_param_1,
	.param .u64 regionaddv_param_2,
	.param .u64 regionaddv_param_3,
	.param .u64 regionaddv_param_4,
	.param .u64 regionaddv_param_5,
	.param .u64 regionaddv_param_6,
	.param .u32 regionaddv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<10>;
	.reg .s64 	%rd<27>;


	ld.param.u64 	%rd8, [regionaddv_param_0];
	ld.param.u64 	%rd9, [regionaddv_param_1];
	ld.param.u64 	%rd10, [regionaddv_param_2];
	ld.param.u64 	%rd11, [regionaddv_param_3];
	ld.param.u64 	%rd12, [regionaddv_param_4];
	ld.param.u64 	%rd13, [regionaddv_param_5];
	ld.param.u64 	%rd14, [regionaddv_param_6];
	ld.param.u32 	%r2, [regionaddv_param_7];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd13;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd11;
	cvta.to.global.u64 	%rd7, %rd14;
	.loc 1 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 10 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_2;

	cvt.s64.s32	%rd15, %r1;
	add.s64 	%rd16, %rd7, %rd15;
	.loc 1 12 1
	ld.global.nc.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	cvt.s64.s8 	%rd18, %rd17;
	shl.b64 	%rd19, %rd18, 2;
	add.s64 	%rd20, %rd6, %rd19;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd5, %rd21;
	.loc 1 13 1
	ld.global.f32 	%f1, [%rd22];
	ld.global.nc.f32 	%f2, [%rd20];
	add.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd22], %f3;
	add.s64 	%rd23, %rd4, %rd19;
	add.s64 	%rd24, %rd3, %rd21;
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd24];
	ld.global.nc.f32 	%f5, [%rd23];
	add.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd24], %f6;
	add.s64 	%rd25, %rd2, %rd19;
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 15 1
	ld.global.f32 	%f7, [%rd26];
	ld.global.nc.f32 	%f8, [%rd25];
	add.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd26], %f9;

BB2_2:
	.loc 1 17 2
	ret;
}


`
)
