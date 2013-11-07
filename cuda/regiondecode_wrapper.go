package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var regiondecode_code cu.Function

type regiondecode_args struct {
	arg_dst     unsafe.Pointer
	arg_LUT     unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [4]unsafe.Pointer
}

// Wrapper for regiondecode CUDA kernel, asynchronous.
func k_regiondecode_async(dst unsafe.Pointer, LUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config, str int) {
	if synchronous { // debug
		SyncAll()
	}

	if regiondecode_code == 0 {
		regiondecode_code = fatbinLoad(regiondecode_map, "regiondecode")
	}

	var _a_ regiondecode_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_LUT = LUT
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_LUT)
	_a_.arg_regions = regions
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_regions)
	_a_.arg_N = N
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_N)

	args := _a_.argptr[:]
	cu.LaunchKernel(regiondecode_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)

	if synchronous { // debug
		SyncAll()
	}
}

// Wrapper for regiondecode CUDA kernel, synchronized.
func k_regiondecode(dst unsafe.Pointer, LUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	const stream = 0
	k_regiondecode_async(dst, LUT, regions, N, cfg, stream)
	Sync(stream)
}

var regiondecode_map = map[int]string{0: "",
	20: regiondecode_ptx_20,
	30: regiondecode_ptx_30,
	35: regiondecode_ptx_35}

const (
	regiondecode_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry regiondecode(
	.param .u64 regiondecode_param_0,
	.param .u64 regiondecode_param_1,
	.param .u64 regiondecode_param_2,
	.param .u32 regiondecode_param_3
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd4, [regiondecode_param_0];
	ld.param.u64 	%rd5, [regiondecode_param_1];
	ld.param.u64 	%rd6, [regiondecode_param_2];
	ld.param.u32 	%r2, [regiondecode_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd5;
	cvta.to.global.u64 	%rd3, %rd6;
	.loc 1 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 7 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvt.s64.s32	%rd7, %r1;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.s8 	%rd9, [%rd8];
	shl.b64 	%rd10, %rd9, 2;
	add.s64 	%rd11, %rd2, %rd10;
	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd1, %rd12;
	.loc 1 9 1
	ld.global.f32 	%f1, [%rd11];
	st.global.f32 	[%rd13], %f1;

BB0_2:
	.loc 1 12 2
	ret;
}


`
	regiondecode_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry regiondecode(
	.param .u64 regiondecode_param_0,
	.param .u64 regiondecode_param_1,
	.param .u64 regiondecode_param_2,
	.param .u32 regiondecode_param_3
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd4, [regiondecode_param_0];
	ld.param.u64 	%rd5, [regiondecode_param_1];
	ld.param.u64 	%rd6, [regiondecode_param_2];
	ld.param.u32 	%r2, [regiondecode_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd5;
	cvta.to.global.u64 	%rd3, %rd6;
	.loc 1 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 7 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvt.s64.s32	%rd7, %r1;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.s8 	%rd9, [%rd8];
	shl.b64 	%rd10, %rd9, 2;
	add.s64 	%rd11, %rd2, %rd10;
	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd1, %rd12;
	.loc 1 9 1
	ld.global.f32 	%f1, [%rd11];
	st.global.f32 	[%rd13], %f1;

BB0_2:
	.loc 1 12 2
	ret;
}


`
	regiondecode_ptx_35 = `
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

.visible .entry regiondecode(
	.param .u64 regiondecode_param_0,
	.param .u64 regiondecode_param_1,
	.param .u64 regiondecode_param_2,
	.param .u32 regiondecode_param_3
)
{
	.reg .pred 	%p<2>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd4, [regiondecode_param_0];
	ld.param.u64 	%rd5, [regiondecode_param_1];
	ld.param.u64 	%rd6, [regiondecode_param_2];
	ld.param.u32 	%r2, [regiondecode_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd5;
	cvta.to.global.u64 	%rd3, %rd6;
	.loc 1 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 7 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_2;

	cvt.s64.s32	%rd7, %r1;
	add.s64 	%rd8, %rd3, %rd7;
	.loc 1 9 1
	ld.global.nc.u8 	%rs1, [%rd8];
	cvt.u64.u16	%rd9, %rs1;
	cvt.s64.s8 	%rd10, %rd9;
	shl.b64 	%rd11, %rd10, 2;
	add.s64 	%rd12, %rd2, %rd11;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd1, %rd13;
	.loc 1 9 1
	ld.global.nc.f32 	%f1, [%rd12];
	st.global.f32 	[%rd14], %f1;

BB2_2:
	.loc 1 12 2
	ret;
}


`
)
