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

func TestNextUnset_All(t *testing.T) {
    var from uint64 = 0
    
    excluded := &[]bool { false, false, false, false }

    from = nextIncluded(from, excluded)
    if (from != 0) {
        t.Fail()
    }

    from = nextIncluded(from + 1, excluded)
    if (from != 1) {
        t.Fail()
    }

    from = nextIncluded(from + 1, excluded)
    if (from != 2) {
        t.Fail()
    }

    from = nextIncluded(from + 1, excluded)
    if (from != 3) {
        t.Fail()
    }
}

func TestNextUnset_Skipped(t *testing.T) {
    var from uint64 = 0
    
    excluded := &[]bool { false, true, false, false }

    from = nextIncluded(from, excluded)
    if (from != 0) {
        t.Fail()
    }

    from = nextIncluded(from + 1, excluded)
    if (from != 2) {
        t.Fail()
    }

    from = nextIncluded(from + 1, excluded)
    if (from != 3) {
        t.Fail()
    }
}


func TestNextUnset_SkippedAll(t *testing.T) {
    var from uint64 = 0
    
    excluded := &[]bool { true, true, true, true }

    from = nextIncluded(from, excluded)
    if (from != 4) {
        t.Fail()
    }
}

func TestFairFast(t *testing.T) {
    var totalBuckets uint64 = Fact(4)    
    excluded := &[]bool { false, false, false, false }
    k_0_4 := fairfast(0, 4, excluded, 0, totalBuckets)
    k_1_4 := fairfast(1, 4, excluded, 0, totalBuckets)
    k_7_4 := fairfast(7, 4, excluded, 0, totalBuckets)
    k_8_4 := fairfast(8, 4, excluded, 0, totalBuckets)

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