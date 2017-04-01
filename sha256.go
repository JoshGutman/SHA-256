package main


import "fmt"
import "strconv"

func SHA256(message string) (string) {
    return "placeholder"
}


func preprocessing(message string) (m []string) {
    bin := padMessage(message)
    m = parseMessage(bin)
    return
}


func padMessage(message string) (string) {

    bin := stringToBin(message)
    bin += "1"
    
    var i int
    for i = 0; len(bin)%512 != 448; i++ {
        bin += "0"
    }
    
    bin = fmt.Sprintf("%s%.64b", bin, i)
    
    return bin
}


func stringToBin(str string) (out string) {
    for _, c := range str {
        out = fmt.Sprintf("%s%.8b", out, c);
    }
    return
}


func parseMessage(message string) (m []string) {
    for i := 0; i <= len(message)-512; i += 512 {
        m = append(m, message[i:i+512])
    }
    return
}


func ch(x, y, z uint32) (uint32) {
    return (x & y) ^ ((^x) & z)
}

func maj(x, y, z uint32) (uint32) {
    return (x & y) ^ (x & z) ^ (y & z)
}

func rotrs(x, n uint32) (uint32) {
    temp := ""   
    
    bin := fmt.Sprintf("%.32b", x)
    
    temp += bin[uint32(len(bin))-n:]
    temp += bin[:uint32(len(bin))-n]
    
    out, _ := strconv.ParseInt(temp, 2, 32)
    return uint32(out)
}

// Applicable for 32 bit ints only
func rotr(x uint32, n int) (uint32) {
    for i := 0; i < n; i++ {
        x = ((x&1)*2147483648) + x>>1
    }
    return x
}

func bsigma0 (x uint32) (uint32) {
    return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func bsigma1 (x uint32) (uint32) {
    return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func lsigma0 (x uint32) (uint32) {
    return rotr(x, 7) ^ rotr(x, 18) ^ (x >> 3)
}

func lsigma1 (x uint32) (uint32) {
    return rotr(x, 17) ^ rotr(x, 19) ^ (x >> 10)
}


func main() {
    //fmt.Printf("%s\n", padMessage("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
    a := preprocessing("a")
    fmt.Printf("%v\n", a)
    fmt.Printf("%d\n", rotr(240, 4))
}
