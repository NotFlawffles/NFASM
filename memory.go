package main

const (
    MemorySize        = 4096
    SegmentTextStart  = 0
    SegmentDataStart  = 1024
    SegmentHeapStart  = 2048
    SegmentStackStart = MemorySize

    SegmentTextSize = SegmentDataStart - SegmentTextStart
    SegmentDataSize = SegmentHeapStart - SegmentDataStart
)
