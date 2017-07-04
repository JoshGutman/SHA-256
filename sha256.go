package main


import "fmt"
import "strconv"
import "os"


func SHA256(message string) (string) {

    m := preprocessing(message)

	var h0, h1, h2, h3, h4, h5, h6, h7 uint32
	h0 = 0x6a09e667
	h1 = 0xbb67ae85
	h2 = 0x3c6ef372
	h3 = 0xa54ff53a
	h4 = 0x510e527f
	h5 = 0x9b05688c
	h6 = 0x1f83d9ab
	h7 = 0x5be0cd19

	k := [64]uint32 {
		0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}
	
	
    for _, block := range(m) {
		
		var w [64]uint32

		for t := 0; t < 16; t++ {
			x, _ := strconv.ParseInt(block[t*32:(t*32)+32], 2, 64)
			w[t] = uint32(x)
		}
		
        for t := 16; t < 64; t++ {
            w[t] = lsigma1(w[t-2]) + w[t-7] + lsigma0(w[t-15]) + w[t-16]
		}

		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		f := h5
		g := h6
		h := h7

		for t := 0; t < 64; t += 1 {
			var t1, t2 uint32
			t1 = h + bsigma1(e) + ch(e, f, g) + k[t] + w[t]
			t2 = bsigma0(a) + maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + t1
			d = c
			c = b
			b = a
			a = t1 + t2
		}

		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
		h5 += f
		h6 += g
		h7 += h
    }

	var out string
	out = fmt.Sprintf("%.8x", h0) + fmt.Sprintf("%.8x", h1) + fmt.Sprintf("%.8x", h2) + fmt.Sprintf("%.8x", h3) + fmt.Sprintf("%.8x", h4) + fmt.Sprintf("%.8x", h5) + fmt.Sprintf("%.8x", h6) + fmt.Sprintf("%.8x", h7)
	return out
}


func preprocessing(message string) (m []string) {
    bin := padMessage(message)
    m = parseMessage(bin)
    return
}


func padMessage(message string) (string) {

	//TODO - make bin a byte array,  When appending 1, just append 10000000 (128) to the last spot in bin
	bin := stringToBin(message)
    bin += "1"
    
    var i int
    for i = 0; len(bin)%512 != 448; i++ {
        bin += "0"
    }
    
	bin = fmt.Sprintf("%s%.64b", bin, len(message)*8)
	
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
    args := os.Args
	for i := 1; i < len(args); i++ {
		fmt.Println(SHA256(args[i]))
	}
}
