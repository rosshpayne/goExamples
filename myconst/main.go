package main

import (
        "fmt"
       )

type Flags uint
type Int   int

const ( Up          Int =  iota
        Broadcast
        LoopBack   
        PointtoPoint
        Multicast
     )
const ( FlagsUp  Flags  = 1 << iota
        FlagsBroadcast
        FlagsLoopBack   = 8 << iota
        FlagsPointtoPoint
        FlagsMulticast
     )
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

type Num int

const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)

func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

func main () {
     fmt.Printf("%d  %b\n", Up, Up)
     fmt.Printf("%d  %b\n", Broadcast, Broadcast)
     fmt.Printf("%d  %b\n", LoopBack, LoopBack)
     fmt.Printf("%d  %b\n", PointtoPoint, PointtoPoint)
     fmt.Printf("%d  %b\n", Multicast, Multicast)
     fmt.Printf("%d  %b\n", FlagsUp, FlagsUp)
     fmt.Printf("%d  %b\n", FlagsBroadcast, FlagsBroadcast)
     fmt.Printf("%d  %b\n", FlagsLoopBack, FlagsLoopBack)
     fmt.Printf("%d  %b\n", FlagsPointtoPoint, FlagsPointtoPoint)
     fmt.Printf("%d  %b\n", FlagsMulticast, FlagsMulticast)
	
     fmt.Printf("\n Placebo: %d",Placebo)
     fmt.Printf("\n Aspirin: %d",Aspirin)
     fmt.Printf("\n Ibuprofen: %d",Ibuprofen)
     fmt.Printf("\n Paracetamol: %d",Paracetamol)
     fmt.Printf("\n Acetaminophenlacebo: %d",Acetaminophen)

     fmt.Printf("\n m_2 = %d",m_2)
     fmt.Printf("\n m_1 = %d",m_1)
     fmt.Printf("\n m0 = %d",m0)
     fmt.Printf("\n m1 = %d",m1)
     fmt.Printf("\n m2 = %d",m2)


     fmt.Printf("\n\n FlagsUp       %d %08[1]b",FlagsUp)
        fmt.Printf("\n FlagsBroadcast     %d %08[1]b",FlagsBroadcast)
        fmt.Printf("\n FlagsLoopBack       %d %08[1]b",FlagsLoopBack)
        fmt.Printf("\n FlagsPointtoPoint  %d %08[1]b",FlagsPointtoPoint)
        fmt.Printf("\n FlagsMulticast     %d %08[1]b",FlagsMulticast)

     x:=KB;fmt.Printf("\n\n x.String() %s",x)
     x=MB;fmt.Printf("\n x.String() %s",x)
     x=GB;fmt.Printf("\n x.String() %s",x)
     x=TB;fmt.Printf("\n x.String() %s",x)
     x=PB;fmt.Printf("\n x.String() %s",x)
     x=ZB;fmt.Printf("\n x.String() %s",x)
     x=YB;fmt.Printf("\n x.String() %s",x)

}


