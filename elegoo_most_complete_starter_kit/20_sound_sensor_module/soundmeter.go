package main

type SoundMeter struct {
	baseLevel uint16
	maxLevel  uint16
	levels    [8]uint16
}

func (sm *SoundMeter) AddSample(level uint16) {
	// Shift all samples down 1, losing oldest (index 0) and adding newest (last index)
	for i := 1; i < len(sm.levels); i++ {
		sm.levels[i-1] = sm.levels[i]
	}
	sm.levels[len(sm.levels)-1] = level
}

func (sm *SoundMeter) Display() {
	for i := 0; i < len(sm.levels); i++ {
		level := sm.levels[len(sm.levels)-1-i]
		// Map sound sensor level to display level
		rowVal := 0b11111111 // 0xFF - Assume 'full' to start
		// Loop over 8 thresholds shifting rowVal right until level > threshold
		smRange := sm.maxLevel - sm.baseLevel + 1
		bi := uint16(0)
		for ; bi < 9; bi++ {
			threshold := sm.maxLevel - smRange*bi/8
			if level > threshold {
				break
			}
		}
		setLEDRow(i, byte((rowVal>>bi)&0xFF))
	}
}
