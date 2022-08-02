package etl

import (
	"strconv"
	"strings"
)

type weekDay int

type schedule struct {
	fromWeekday *weekDay
	toWeekDay   *weekDay
	from        *clock
	to          *clock
}

type clock struct {
	hour   int
	minute int
}

var m = map[string]weekDay{
	"Sun":   0,
	"Mon":   1,
	"Tues":  2,
	"Weds":  3,
	"Thurs": 4,
	"Fri":   5,
	"Sat":   6,
}

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func (w *weekDay) String() string {
	return longDayNames[*w]
}

func parseOpeningHours(openHours string) []*schedule {
	days := strings.Split(openHours, "/")
	schedules := make([]*schedule, 0)

	for i := range days {
		if days[i] == "" {
			continue
		}
		schedules = append(schedules, parseSchedule(days[i]))
	}
	return schedules
}

func parseSchedule(openHour string) *schedule {
	s := &schedule{}
	s.parseOpeningHour(openHour)
	return s
}

func (s *schedule) parseOpeningHour(openHour string) {
	i := 0
	// Mon - Tues 7:45 am - 11:15 am
	// Tues 1:30 pm - 2:30 pm
	// Thurs, Sun 8 am - 1:15 am
	dashFound := false
	commaFound := false
	weekDayShouldFind := true
	for ; i < len(openHour); i++ {
		if !dashFound && !commaFound && weekDayShouldFind && isNumber(rune(openHour[i+2])) {
			s.setFromWeekDay(openHour[:i+1])
			s.setToWeekDay("")
			weekDayShouldFind = false
		}

		switch openHour[i] {
		case '-':
			if openHour[i-3:i-1] == "pm" || openHour[i-3:i-1] == "am" {
				in := i - 8
				for ; in <= i-3; in++ {
					if !isNumber(rune(openHour[in])) {
						continue
					}
					break
				}
				s.setFromTime(openHour[in : i-1])
				s.setToTime(openHour[i+1:])
				return
			} else {
				s.setFromWeekDay(openHour[:i-1])
				s.setToWeekDay(openHour[i+1 : i+6])
				dashFound = true
			}
		case ',':
			s.setFromWeekDay(openHour[:i])
			s.setToWeekDay(openHour[i+2 : i+5])
			commaFound = true
		}
	}
}

func (s *schedule) setFromWeekDay(t string) {
	w := m[strings.Trim(t, " ")]
	s.fromWeekday = &w
}

func (s *schedule) setToWeekDay(t string) {
	if t == "" {
		s.toWeekDay = nil
		return
	}
	w := m[strings.Trim(t, " ")]
	s.toWeekDay = &w
}

func (s *schedule) setFromTime(t string) {
	h, r := parseTime(strings.Trim(t, " "))
	s.from = &clock{
		hour:   h,
		minute: r,
	}
}

func (s *schedule) setToTime(t string) {
	h, r := parseTime(strings.Trim(t, " "))
	s.to = &clock{
		hour:   h,
		minute: r,
	}
}

func parseTime(t string) (int, int) {
	ami := strings.Index(t, "am")
	pmi := strings.Index(t, "pm")
	if ami > -1 {
		if len(t[:ami-1]) > 0 {
			return strToHourMin(t[:ami-1])
		}
	}

	if pmi > -1 {
		if len(t[:pmi-1]) > 0 {
			return convertTo24(t[:pmi-1])
		}
	}

	return 0, 0
}

func strToHourMin(s string) (int, int) {
	sp := strings.Split(s, ":")

	if len(sp) < 1 {
		return 0, 0
	}

	h, err := strconv.Atoi(sp[0])
	if err != nil {
		return 0, 0
	}

	if len(sp) < 2 {
		return h, 0
	}

	m, err := strconv.Atoi(sp[1])
	if err != nil {
		return 0, 0
	}
	return h, m
}

func convertTo24(s string) (int, int) {
	h, m := strToHourMin(s)
	return h + 12, m
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
