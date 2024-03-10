package MovingAverageFromDataStream

type MovingAverage struct {
	Window      int
	WindowTotal int
	Stream      []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Window: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.Stream) < this.Window {
		this.Stream = append(this.Stream, val)
	} else {
		this.WindowTotal -= this.Stream[0]
		this.Stream = append(this.Stream[1:], val)
	}
	this.WindowTotal += val
	return float64(this.WindowTotal) / float64(len(this.Stream))
}
