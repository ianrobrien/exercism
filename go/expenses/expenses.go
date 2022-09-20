package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record

	for i := range in {
		if predicate(in[i]) {
			out = append(out, in[i])
		}
	}

	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	amount := 0.0
	for i := range in {
		if (ByDaysPeriod(p))(in[i]) {
			amount += in[i].Amount
		}
	}
	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	expenses := 0.0
	var hasCategory bool

	for i := range in {
		if ByCategory(c)(in[i]) {
			hasCategory = true
			if ByDaysPeriod(p)(in[i]) {
				expenses += in[i].Amount
			}
		}
	}

	if !hasCategory {
		return 0, errors.New("no expenses for given category")
	}

	return expenses, nil
}
