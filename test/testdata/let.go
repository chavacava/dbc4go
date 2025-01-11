package testdata

// @let foo := 1
// @ensures foo == 1
func letRaw() {}

// Contract:
//   - let foo := 1
//   - ensures foo == 1
func letStandard() {}
