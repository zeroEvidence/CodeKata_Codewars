package kata_ball_upwards

func MaxBall(v0 int) int {
  var i int = 0
  var h float32 = 0
  var v float32 = float32(v0) / 3.6
  var t float32 = 0.00
  var g float32 = 9.81

  for {

    heightNow := (v*t) - (0.5*g*t*t)

    if heightNow < h {
      break
    } else {
      h = heightNow

      i++
      t = float32(i) * 0.1
    }
  }

  return i - 1
}