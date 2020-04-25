package cryptorandomizer        //This package make randoom nums for you
import (
    crand "crypto/rand"
    rand "math/rand"
    "encoding/binary"
    "log"
)
                                // Num is awesome
func Num(x int) int{            // func Num function using input as limiter to how big is max
    var src cryptoSource        // random number and returning it as integer
    rnd := rand.New(src)        // why is he still complaining?
    //fmt.Println(rnd.Intn(x))  // what do you want from me bear?
    return rnd.Intn(x)
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
    return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
    err := binary.Read(crand.Reader, binary.BigEndian, &v)
    if err != nil {
        log.Fatal(err)
    }
    return v
}
