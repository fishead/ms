package ms

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	if _, err := Parse("1m"); err != nil {
		t.Error("should not throw an error")
	}

	if ms, err := Parse("100"); err != nil || ms != 100 {
		fmt.Println(err)
		t.Error("does not preserve ms")
	}

	if ms, err := Parse("1m"); err != nil || ms != 60000 {
		t.Error("does not convert from m to ms")
	}

	if ms, err := Parse("1h"); err != nil || ms != 3600000 {
		t.Error("does not convert from h to ms")
	}

	if ms, err := Parse("2d"); err != nil || ms != 172800000 {
		fmt.Println(ms)
		t.Error("does not convert d to ms")
	}

	if ms, err := Parse("3w"); err != nil || ms != 1814400000 {
		t.Error("does not convert w to ms")
	}

	if ms, err := Parse("1s"); err != nil || ms != 1000 {
		t.Error("does not convert s to ms")
	}

	if ms, err := Parse("100ms"); err != nil || ms != 100 {
		t.Error("does not convert ms to ms")
	}

	if ms, err := Parse("1.5h"); err != nil || ms != 5400000 {
		t.Error("does not work with decimals")
	}

	if ms, err := Parse("1   s"); err != nil || ms != 1000 {
		t.Error("does not work with multiple spaces")
	}

	if _, err := Parse("☃"); err == nil {
		t.Error("does not return error if invalid")
	}

	if ms, err := Parse("1.5H"); err != nil || ms != 5400000 {
		t.Error("should not be case-sensitive")
	}

	if ms, err := Parse(".5ms"); err != nil || ms != 0.5 {
		t.Error("does not work with numbers starting with .")
	}

	if ms, err := Parse("-100ms"); err != nil || ms != -100 {
		t.Error("does not work with negative integers")
	}

	if ms, err := Parse("-1.5h"); err != nil || ms != -5400000 {
		t.Error("does not work with negative decimals")
	}

	if ms, err := Parse("-.5h"); err != nil || ms != -1800000 {
		t.Error(`does not work with negative decimals starting with "."`)
	}
}

// ms(long string)
func Test2(t *testing.T) {
	if _, err := Parse("53 milliseconds"); err != nil {
		fmt.Println(err)
		t.Error("should not throw an error")
	}

	if ms, _ := Parse("53 milliseconds"); ms != 53 {
		fmt.Println(ms)
		t.Error("does not convert milliseconds to ms")
	}

	if ms, _ := Parse("17 msecs"); ms != 17 {
		t.Error("should convert msecs to ms")
	}

	if ms, _ := Parse("1 sec"); ms != 1000 {
		t.Error("should convert sec to ms")
	}

	if ms, _ := Parse("1 min"); ms != 60000 {
		t.Error("should convert from min to ms")
	}

	if ms, _ := Parse("1 hr"); ms != 3600000 {
		t.Error("should convert from hr to ms")
	}

	if ms, _ := Parse("2 days"); ms != 172800000 {
		t.Error("should convert days to ms")
	}

	if ms, _ := Parse("1.5 hours"); ms != 5400000 {
		t.Error("should work with decimals")
	}

	if ms, _ := Parse("-100 milliseconds"); ms != -100 {
		t.Error("should work with negative integers")
	}

	if ms, _ := Parse("-1.5 hours"); ms != -5400000 {
		t.Error("should work with negative decimals")
	}

	if ms, _ := Parse("-.5 hr"); ms != -1800000 {
		t.Error(`should work with negative decimals starting with "."`)
	}
}

// Test3 FormatLong(number, { long: true })
func Test3(t *testing.T) {
	if ms := FmtLong(500); ms != "500 ms" {
		fmt.Println(ms, "11111")
		t.Error("should support milliseconds")
	}

	if ms := FmtLong(-500); ms != "-500 ms" {
		t.Error("should support negative milliseconds")
	}

	if ms := FmtLong(1000); ms != "1 second" {
		t.Error("should support seconds")

	}
	if ms := FmtLong(1200); ms != "1 second" {
		t.Error("should support seconds")

	}
	if ms := FmtLong(10000); ms != "10 seconds" {
		t.Error("should support seconds")

	}

	if ms := FmtLong(-1000); ms != "-1 second" {
		t.Error("should support seconds")

	}
	if ms := FmtLong(-1200); ms != "-1 second" {
		t.Error("should support seconds")

	}
	if ms := FmtLong(-10000); ms != "-10 seconds" {
		t.Error("should support seconds")

	}

	if ms := FmtLong(60 * 1000); ms != "1 minute" {
		t.Error("should support minutes")
	}
	if ms := FmtLong(60 * 1200); ms != "1 minute" {
		t.Error("should support minutes")
	}
	if ms := FmtLong(60 * 10000); ms != "10 minutes" {
		t.Error("should support minutes")
	}

	if ms := FmtLong(-1 * 60 * 1000); ms != "-1 minute" {
		t.Error("should support minutes")
	}
	if ms := FmtLong(-1 * 60 * 1200); ms != "-1 minute" {
		t.Error("should support minutes")
	}
	if ms := FmtLong(-1 * 60 * 10000); ms != "-10 minutes" {
		t.Error("should support minutes")
	}

	if ms := FmtLong(60 * 60 * 1000); ms != "1 hour" {
		t.Error("should support hours")
	}
	if ms := FmtLong(60 * 60 * 1200); ms != "1 hour" {
		t.Error("should support hours")
	}
	if ms := FmtLong(60 * 60 * 10000); ms != "10 hours" {
		t.Error("should support hours")
	}

	if ms := FmtLong(-1 * 60 * 60 * 1000); ms != "-1 hour" {
		t.Error("should support hours")
	}
	if ms := FmtLong(-1 * 60 * 60 * 1200); ms != "-1 hour" {
		t.Error("should support hours")
	}
	if ms := FmtLong(-1 * 60 * 60 * 10000); ms != "-10 hours" {
		t.Error("should support hours")
	}

	if ms := FmtLong(24 * 60 * 60 * 1000); ms != "1 day" {
		t.Error("should support days")
	}
	if ms := FmtLong(24 * 60 * 60 * 1200); ms != "1 day" {
		t.Error("should support days")
	}
	if ms := FmtLong(24 * 60 * 60 * 10000); ms != "10 days" {
		t.Error("should support days")
	}

	if ms := FmtLong(-1 * 24 * 60 * 60 * 1000); ms != "-1 day" {
		t.Error("should support days")
	}
	if ms := FmtLong(-1 * 24 * 60 * 60 * 1200); ms != "-1 day" {
		t.Error("should support days")
	}
	if ms := FmtLong(-1 * 24 * 60 * 60 * 10000); ms != "-10 days" {
		t.Error("should support days")
	}

	if ms := FmtLong(234234234); ms != "3 days" {
		t.Error("should round")
	}

	if ms := FmtLong(-234234234); ms != "-3 days" {
		t.Error("should round")
	}
}

// Test4 FmtShort
func Test4(t *testing.T) {
	if ms := FmtShort(500); ms != "500ms" {
		t.Error("should support milliseconds")
	}

	if ms := FmtShort(-500); ms != "-500ms" {
		t.Error("should support milliseconds")
	}

	if ms := FmtShort(1000); ms != "1s" {
		t.Error("should support seconds")
	}
	if ms := FmtShort(10000); ms != "10s" {
		t.Error("should support seconds")
	}

	if ms := FmtShort(-1000); ms != "-1s" {
		t.Error("should support seconds")
	}
	if ms := FmtShort(-10000); ms != "-10s" {
		t.Error("should support seconds")
	}

	if ms := FmtShort(60 * 1000); ms != "1m" {
		t.Error("should support minutes")
	}
	if ms := FmtShort(60 * 10000); ms != "10m" {
		t.Error("should support minutes")
	}

	if ms := FmtShort(-1 * 60 * 1000); ms != "-1m" {
		t.Error("should support minutes")
	}
	if ms := FmtShort(-1 * 60 * 10000); ms != "-10m" {
		t.Error("should support minutes")
	}

	if ms := FmtShort(60 * 60 * 1000); ms != "1h" {
		t.Error("should support hours")
	}
	if ms := FmtShort(60 * 60 * 10000); ms != "10h" {
		t.Error("should support hours")
	}

	if ms := FmtShort(-1 * 60 * 60 * 1000); ms != "-1h" {
		t.Error("should support hours")
	}
	if ms := FmtShort(-1 * 60 * 60 * 10000); ms != "-10h" {
		t.Error("should support hours")
	}

	if ms := FmtShort(24 * 60 * 60 * 1000); ms != "1d" {
		t.Error("should support days")
	}
	if ms := FmtShort(24 * 60 * 60 * 10000); ms != "10d" {
		t.Error("should support days")
	}

	if ms := FmtShort(-1 * 24 * 60 * 60 * 1000); ms != "-1d" {
		t.Error("should support days")
	}
	if ms := FmtShort(-1 * 24 * 60 * 60 * 10000); ms != "-10d" {
		t.Error("should support days")
	}

	if ms := FmtShort(234234234); ms != "3d" {
		t.Error("should round")
	}

	if ms := FmtShort(-234234234); ms != "-3d" {
		t.Error("should round")
	}
}
