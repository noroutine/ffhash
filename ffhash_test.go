package ffhash

import (
    "testing"
    "flag"
    "os"
)

func TestMain(m *testing.M) {
    flag.Parse()
    os.Exit(m.Run())
}

func TestFairFast(t *testing.T) {
    var totalBuckets uint64 = Fact(4)
    k_0_4 := fairfast(0, 4, 0, totalBuckets)
    k_1_4 := fairfast(1, 4, 0, totalBuckets)
    k_7_4 := fairfast(7, 4, 0, totalBuckets)
    k_8_4 := fairfast(8, 4, 0, totalBuckets)

    if k_0_4 != 3 {
        t.Fail()
    }

    if k_1_4 != 2 {
        t.Fail()
    }

    if k_7_4 != 1 {
        t.Fail()
    }

    if k_8_4 != 3 {
        t.Fail()
    }
}