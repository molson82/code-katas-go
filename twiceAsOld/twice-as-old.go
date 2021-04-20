package twiceAsOld

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
  if sonYearsOld == 0 {
    return dadYearsOld
  }

  return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}
