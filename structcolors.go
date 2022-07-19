package main


type colors struct {
  Reset string
  Bold string
  Red string
  Green string
  Yellow string
  Blue string
  Cyan string
  White string
}

func (c *colors) DisableColors() {
  c.Reset = ""
  c.Bold = ""
  c.Red = ""
  c.Green = ""
  c.Yellow = ""
  c.Blue = ""
  c.Cyan = ""
  c.White = ""
}


var Color = colors {
    Reset:    "\033[m",
    Bold:     "\033[1m",

    Red:    "\033[31m",
    Green:  "\033[32m",
    Yellow: "\033[33m",
    Blue:   "\033[34m",
    Cyan:   "\033[36m",
    White:  "\033[37m",
}
