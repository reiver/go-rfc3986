package rfc3986_test

import (
	"testing"

	"math/rand"
	"time"

	"github.com/reiver/go-rfc3986"
)

func TestHasPrefixPctEncoded(t *testing.T) {
	tests := []struct{
		String string
		Expected bool
	}{
		{
			String: "",
			Expected: false,
		},



		{
			String: "ONCE",
			Expected: false,
		},
		{
			String: "TWICE",
			Expected: false,
		},
		{
			String: "THRICE",
			Expected: false,
		},
		{
			String: "FOURCE",
			Expected: false,
		},



		{
			String: "0x00",
			Expected: false,
		},
		{
			String: "0xff",
			Expected: false,
		},
		{
			String: "0xFF",
			Expected: false,
		},



		{
			String: "%",
			Expected: false,
		},



		{
			String: "%GG",
			Expected: false,
		},
		{
			String: "%Gg",
			Expected: false,
		},
		{
			String: "%gG",
			Expected: false,
		},
		{
			String: "%gg",
			Expected: false,
		},



		{
			String: "%0",
			Expected: false,
		},
		{
			String: "%1",
			Expected: false,
		},
		{
			String: "%2",
			Expected: false,
		},
		{
			String: "%3",
			Expected: false,
		},
		{
			String: "%4",
			Expected: false,
		},
		{
			String: "%5",
			Expected: false,
		},
		{
			String: "%6",
			Expected: false,
		},
		{
			String: "%7",
			Expected: false,
		},
		{
			String: "%8",
			Expected: false,
		},
		{
			String: "%9",
			Expected: false,
		},
		{
			String: "%A",
			Expected: false,
		},
		{
			String: "%a",
			Expected: false,
		},
		{
			String: "%B",
			Expected: false,
		},
		{
			String: "%b",
			Expected: false,
		},
		{
			String: "%C",
			Expected: false,
		},
		{
			String: "%c",
			Expected: false,
		},
		{
			String: "%D",
			Expected: false,
		},
		{
			String: "%d",
			Expected: false,
		},
		{
			String: "%E",
			Expected: false,
		},
		{
			String: "%e",
			Expected: false,
		},
		{
			String: "%F",
			Expected: false,
		},
		{
			String: "%f",
			Expected: false,
		},



		{
			String: "%0G",
			Expected: false,
		},
		{
			String: "%1G",
			Expected: false,
		},
		{
			String: "%2G",
			Expected: false,
		},
		{
			String: "%3G",
			Expected: false,
		},
		{
			String: "%4G",
			Expected: false,
		},
		{
			String: "%5G",
			Expected: false,
		},
		{
			String: "%6G",
			Expected: false,
		},
		{
			String: "%7G",
			Expected: false,
		},
		{
			String: "%8G",
			Expected: false,
		},
		{
			String: "%9G",
			Expected: false,
		},
		{
			String: "%AG",
			Expected: false,
		},
		{
			String: "%aG",
			Expected: false,
		},
		{
			String: "%BG",
			Expected: false,
		},
		{
			String: "%bG",
			Expected: false,
		},
		{
			String: "%CG",
			Expected: false,
		},
		{
			String: "%cG",
			Expected: false,
		},
		{
			String: "%DG",
			Expected: false,
		},
		{
			String: "%dG",
			Expected: false,
		},
		{
			String: "%EG",
			Expected: false,
		},
		{
			String: "%eG",
			Expected: false,
		},
		{
			String: "%FG",
			Expected: false,
		},
		{
			String: "%fG",
			Expected: false,
		},



		{
			String: "%0g",
			Expected: false,
		},
		{
			String: "%1g",
			Expected: false,
		},
		{
			String: "%2g",
			Expected: false,
		},
		{
			String: "%3g",
			Expected: false,
		},
		{
			String: "%4g",
			Expected: false,
		},
		{
			String: "%5g",
			Expected: false,
		},
		{
			String: "%6g",
			Expected: false,
		},
		{
			String: "%7g",
			Expected: false,
		},
		{
			String: "%8g",
			Expected: false,
		},
		{
			String: "%9g",
			Expected: false,
		},
		{
			String: "%Ag",
			Expected: false,
		},
		{
			String: "%ag",
			Expected: false,
		},
		{
			String: "%Bg",
			Expected: false,
		},
		{
			String: "%bg",
			Expected: false,
		},
		{
			String: "%Cg",
			Expected: false,
		},
		{
			String: "%cg",
			Expected: false,
		},
		{
			String: "%Dg",
			Expected: false,
		},
		{
			String: "%dg",
			Expected: false,
		},
		{
			String: "%Eg",
			Expected: false,
		},
		{
			String: "%eg",
			Expected: false,
		},
		{
			String: "%Fg",
			Expected: false,
		},
		{
			String: "%fg",
			Expected: false,
		},



		{
			String: "%00",
			Expected: true,
		},
		{
			String: "%01",
			Expected: true,
		},
		{
			String: "%02",
			Expected: true,
		},
		{
			String: "%03",
			Expected: true,
		},
		{
			String: "%04",
			Expected: true,
		},
		{
			String: "%05",
			Expected: true,
		},
		{
			String: "%06",
			Expected: true,
		},
		{
			String: "%07",
			Expected: true,
		},
		{
			String: "%08",
			Expected: true,
		},
		{
			String: "%09",
			Expected: true,
		},
		{
			String: "%0A",
			Expected: true,
		},
		{
			String: "%0a",
			Expected: false,
		},
		{
			String: "%0B",
			Expected: true,
		},
		{
			String: "%0b",
			Expected: false,
		},
		{
			String: "%0C",
			Expected: true,
		},
		{
			String: "%0c",
			Expected: false,
		},
		{
			String: "%0D",
			Expected: true,
		},
		{
			String: "%0d",
			Expected: false,
		},
		{
			String: "%0E",
			Expected: true,
		},
		{
			String: "%0e",
			Expected: false,
		},
		{
			String: "%0F",
			Expected: true,
		},
		{
			String: "%0f",
			Expected: false,
		},



		{
			String: "%10",
			Expected: true,
		},
		{
			String: "%11",
			Expected: true,
		},
		{
			String: "%12",
			Expected: true,
		},
		{
			String: "%13",
			Expected: true,
		},
		{
			String: "%14",
			Expected: true,
		},
		{
			String: "%15",
			Expected: true,
		},
		{
			String: "%16",
			Expected: true,
		},
		{
			String: "%17",
			Expected: true,
		},
		{
			String: "%18",
			Expected: true,
		},
		{
			String: "%19",
			Expected: true,
		},
		{
			String: "%1A",
			Expected: true,
		},
		{
			String: "%1a",
			Expected: false,
		},
		{
			String: "%1B",
			Expected: true,
		},
		{
			String: "%1b",
			Expected: false,
		},
		{
			String: "%1C",
			Expected: true,
		},
		{
			String: "%1c",
			Expected: false,
		},
		{
			String: "%1D",
			Expected: true,
		},
		{
			String: "%1d",
			Expected: false,
		},
		{
			String: "%1E",
			Expected: true,
		},
		{
			String: "%1e",
			Expected: false,
		},
		{
			String: "%1F",
			Expected: true,
		},
		{
			String: "%1f",
			Expected: false,
		},



		{
			String: "%20",
			Expected: true,
		},
		{
			String: "%21",
			Expected: true,
		},
		{
			String: "%22",
			Expected: true,
		},
		{
			String: "%23",
			Expected: true,
		},
		{
			String: "%24",
			Expected: true,
		},
		{
			String: "%25",
			Expected: true,
		},
		{
			String: "%26",
			Expected: true,
		},
		{
			String: "%27",
			Expected: true,
		},
		{
			String: "%28",
			Expected: true,
		},
		{
			String: "%29",
			Expected: true,
		},
		{
			String: "%2A",
			Expected: true,
		},
		{
			String: "%2a",
			Expected: false,
		},
		{
			String: "%2B",
			Expected: true,
		},
		{
			String: "%2b",
			Expected: false,
		},
		{
			String: "%2C",
			Expected: true,
		},
		{
			String: "%2c",
			Expected: false,
		},
		{
			String: "%2D",
			Expected: true,
		},
		{
			String: "%2d",
			Expected: false,
		},
		{
			String: "%2E",
			Expected: true,
		},
		{
			String: "%2e",
			Expected: false,
		},
		{
			String: "%2F",
			Expected: true,
		},
		{
			String: "%2f",
			Expected: false,
		},



		{
			String: "%30",
			Expected: true,
		},
		{
			String: "%31",
			Expected: true,
		},
		{
			String: "%32",
			Expected: true,
		},
		{
			String: "%33",
			Expected: true,
		},
		{
			String: "%34",
			Expected: true,
		},
		{
			String: "%35",
			Expected: true,
		},
		{
			String: "%36",
			Expected: true,
		},
		{
			String: "%37",
			Expected: true,
		},
		{
			String: "%38",
			Expected: true,
		},
		{
			String: "%39",
			Expected: true,
		},
		{
			String: "%3A",
			Expected: true,
		},
		{
			String: "%3a",
			Expected: false,
		},
		{
			String: "%3B",
			Expected: true,
		},
		{
			String: "%3b",
			Expected: false,
		},
		{
			String: "%3C",
			Expected: true,
		},
		{
			String: "%3c",
			Expected: false,
		},
		{
			String: "%3D",
			Expected: true,
		},
		{
			String: "%3d",
			Expected: false,
		},
		{
			String: "%3E",
			Expected: true,
		},
		{
			String: "%3e",
			Expected: false,
		},
		{
			String: "%3F",
			Expected: true,
		},
		{
			String: "%3f",
			Expected: false,
		},



		{
			String: "%40",
			Expected: true,
		},
		{
			String: "%41",
			Expected: true,
		},
		{
			String: "%42",
			Expected: true,
		},
		{
			String: "%43",
			Expected: true,
		},
		{
			String: "%44",
			Expected: true,
		},
		{
			String: "%45",
			Expected: true,
		},
		{
			String: "%46",
			Expected: true,
		},
		{
			String: "%47",
			Expected: true,
		},
		{
			String: "%48",
			Expected: true,
		},
		{
			String: "%49",
			Expected: true,
		},
		{
			String: "%4A",
			Expected: true,
		},
		{
			String: "%4a",
			Expected: false,
		},
		{
			String: "%4B",
			Expected: true,
		},
		{
			String: "%4b",
			Expected: false,
		},
		{
			String: "%4C",
			Expected: true,
		},
		{
			String: "%4c",
			Expected: false,
		},
		{
			String: "%4D",
			Expected: true,
		},
		{
			String: "%4d",
			Expected: false,
		},
		{
			String: "%4E",
			Expected: true,
		},
		{
			String: "%4e",
			Expected: false,
		},
		{
			String: "%4F",
			Expected: true,
		},
		{
			String: "%4f",
			Expected: false,
		},



		{
			String: "%50",
			Expected: true,
		},
		{
			String: "%51",
			Expected: true,
		},
		{
			String: "%52",
			Expected: true,
		},
		{
			String: "%53",
			Expected: true,
		},
		{
			String: "%54",
			Expected: true,
		},
		{
			String: "%55",
			Expected: true,
		},
		{
			String: "%56",
			Expected: true,
		},
		{
			String: "%57",
			Expected: true,
		},
		{
			String: "%58",
			Expected: true,
		},
		{
			String: "%59",
			Expected: true,
		},
		{
			String: "%5A",
			Expected: true,
		},
		{
			String: "%5a",
			Expected: false,
		},
		{
			String: "%5B",
			Expected: true,
		},
		{
			String: "%5b",
			Expected: false,
		},
		{
			String: "%5C",
			Expected: true,
		},
		{
			String: "%5c",
			Expected: false,
		},
		{
			String: "%5D",
			Expected: true,
		},
		{
			String: "%5d",
			Expected: false,
		},
		{
			String: "%5E",
			Expected: true,
		},
		{
			String: "%5e",
			Expected: false,
		},
		{
			String: "%5F",
			Expected: true,
		},
		{
			String: "%5f",
			Expected: false,
		},



		{
			String: "%60",
			Expected: true,
		},
		{
			String: "%61",
			Expected: true,
		},
		{
			String: "%62",
			Expected: true,
		},
		{
			String: "%63",
			Expected: true,
		},
		{
			String: "%64",
			Expected: true,
		},
		{
			String: "%65",
			Expected: true,
		},
		{
			String: "%66",
			Expected: true,
		},
		{
			String: "%67",
			Expected: true,
		},
		{
			String: "%68",
			Expected: true,
		},
		{
			String: "%69",
			Expected: true,
		},
		{
			String: "%6A",
			Expected: true,
		},
		{
			String: "%6a",
			Expected: false,
		},
		{
			String: "%6B",
			Expected: true,
		},
		{
			String: "%6b",
			Expected: false,
		},
		{
			String: "%6C",
			Expected: true,
		},
		{
			String: "%6c",
			Expected: false,
		},
		{
			String: "%6D",
			Expected: true,
		},
		{
			String: "%6d",
			Expected: false,
		},
		{
			String: "%6E",
			Expected: true,
		},
		{
			String: "%6e",
			Expected: false,
		},
		{
			String: "%6F",
			Expected: true,
		},
		{
			String: "%6f",
			Expected: false,
		},



		{
			String: "%70",
			Expected: true,
		},
		{
			String: "%71",
			Expected: true,
		},
		{
			String: "%72",
			Expected: true,
		},
		{
			String: "%73",
			Expected: true,
		},
		{
			String: "%74",
			Expected: true,
		},
		{
			String: "%75",
			Expected: true,
		},
		{
			String: "%76",
			Expected: true,
		},
		{
			String: "%77",
			Expected: true,
		},
		{
			String: "%78",
			Expected: true,
		},
		{
			String: "%79",
			Expected: true,
		},
		{
			String: "%7A",
			Expected: true,
		},
		{
			String: "%7a",
			Expected: false,
		},
		{
			String: "%7B",
			Expected: true,
		},
		{
			String: "%7b",
			Expected: false,
		},
		{
			String: "%7C",
			Expected: true,
		},
		{
			String: "%7c",
			Expected: false,
		},
		{
			String: "%7D",
			Expected: true,
		},
		{
			String: "%7d",
			Expected: false,
		},
		{
			String: "%7E",
			Expected: true,
		},
		{
			String: "%7e",
			Expected: false,
		},
		{
			String: "%7F",
			Expected: true,
		},
		{
			String: "%7f",
			Expected: false,
		},



		{
			String: "%80",
			Expected: true,
		},
		{
			String: "%81",
			Expected: true,
		},
		{
			String: "%82",
			Expected: true,
		},
		{
			String: "%83",
			Expected: true,
		},
		{
			String: "%84",
			Expected: true,
		},
		{
			String: "%85",
			Expected: true,
		},
		{
			String: "%86",
			Expected: true,
		},
		{
			String: "%87",
			Expected: true,
		},
		{
			String: "%88",
			Expected: true,
		},
		{
			String: "%89",
			Expected: true,
		},
		{
			String: "%8A",
			Expected: true,
		},
		{
			String: "%8a",
			Expected: false,
		},
		{
			String: "%8B",
			Expected: true,
		},
		{
			String: "%8b",
			Expected: false,
		},
		{
			String: "%8C",
			Expected: true,
		},
		{
			String: "%8c",
			Expected: false,
		},
		{
			String: "%8D",
			Expected: true,
		},
		{
			String: "%8d",
			Expected: false,
		},
		{
			String: "%8E",
			Expected: true,
		},
		{
			String: "%8e",
			Expected: false,
		},
		{
			String: "%8F",
			Expected: true,
		},
		{
			String: "%8f",
			Expected: false,
		},



		{
			String: "%90",
			Expected: true,
		},
		{
			String: "%91",
			Expected: true,
		},
		{
			String: "%92",
			Expected: true,
		},
		{
			String: "%93",
			Expected: true,
		},
		{
			String: "%94",
			Expected: true,
		},
		{
			String: "%95",
			Expected: true,
		},
		{
			String: "%96",
			Expected: true,
		},
		{
			String: "%97",
			Expected: true,
		},
		{
			String: "%98",
			Expected: true,
		},
		{
			String: "%99",
			Expected: true,
		},
		{
			String: "%9A",
			Expected: true,
		},
		{
			String: "%9a",
			Expected: false,
		},
		{
			String: "%9B",
			Expected: true,
		},
		{
			String: "%9b",
			Expected: false,
		},
		{
			String: "%9C",
			Expected: true,
		},
		{
			String: "%9c",
			Expected: false,
		},
		{
			String: "%9D",
			Expected: true,
		},
		{
			String: "%9d",
			Expected: false,
		},
		{
			String: "%9E",
			Expected: true,
		},
		{
			String: "%9e",
			Expected: false,
		},
		{
			String: "%9F",
			Expected: true,
		},
		{
			String: "%9f",
			Expected: false,
		},



		{
			String: "%A0",
			Expected: true,
		},
		{
			String: "%A1",
			Expected: true,
		},
		{
			String: "%A2",
			Expected: true,
		},
		{
			String: "%A3",
			Expected: true,
		},
		{
			String: "%A4",
			Expected: true,
		},
		{
			String: "%A5",
			Expected: true,
		},
		{
			String: "%A6",
			Expected: true,
		},
		{
			String: "%A7",
			Expected: true,
		},
		{
			String: "%A8",
			Expected: true,
		},
		{
			String: "%A9",
			Expected: true,
		},
		{
			String: "%AA",
			Expected: true,
		},
		{
			String: "%Aa",
			Expected: false,
		},
		{
			String: "%AB",
			Expected: true,
		},
		{
			String: "%Ab",
			Expected: false,
		},
		{
			String: "%AC",
			Expected: true,
		},
		{
			String: "%Ac",
			Expected: false,
		},
		{
			String: "%AD",
			Expected: true,
		},
		{
			String: "%Ad",
			Expected: false,
		},
		{
			String: "%AE",
			Expected: true,
		},
		{
			String: "%Ae",
			Expected: false,
		},
		{
			String: "%AF",
			Expected: true,
		},
		{
			String: "%Af",
			Expected: false,
		},



		{
			String: "%a0",
			Expected: false,
		},
		{
			String: "%a1",
			Expected: false,
		},
		{
			String: "%a2",
			Expected: false,
		},
		{
			String: "%a3",
			Expected: false,
		},
		{
			String: "%a4",
			Expected: false,
		},
		{
			String: "%a5",
			Expected: false,
		},
		{
			String: "%a6",
			Expected: false,
		},
		{
			String: "%a7",
			Expected: false,
		},
		{
			String: "%a8",
			Expected: false,
		},
		{
			String: "%a9",
			Expected: false,
		},
		{
			String: "%aA",
			Expected: false,
		},
		{
			String: "%aa",
			Expected: false,
		},
		{
			String: "%aB",
			Expected: false,
		},
		{
			String: "%ab",
			Expected: false,
		},
		{
			String: "%aC",
			Expected: false,
		},
		{
			String: "%ac",
			Expected: false,
		},
		{
			String: "%aD",
			Expected: false,
		},
		{
			String: "%ad",
			Expected: false,
		},
		{
			String: "%aE",
			Expected: false,
		},
		{
			String: "%ae",
			Expected: false,
		},
		{
			String: "%aF",
			Expected: false,
		},
		{
			String: "%af",
			Expected: false,
		},



		{
			String: "%B0",
			Expected: true,
		},
		{
			String: "%B1",
			Expected: true,
		},
		{
			String: "%B2",
			Expected: true,
		},
		{
			String: "%B3",
			Expected: true,
		},
		{
			String: "%B4",
			Expected: true,
		},
		{
			String: "%B5",
			Expected: true,
		},
		{
			String: "%B6",
			Expected: true,
		},
		{
			String: "%B7",
			Expected: true,
		},
		{
			String: "%B8",
			Expected: true,
		},
		{
			String: "%B9",
			Expected: true,
		},
		{
			String: "%BA",
			Expected: true,
		},
		{
			String: "%Ba",
			Expected: false,
		},
		{
			String: "%BB",
			Expected: true,
		},
		{
			String: "%Bb",
			Expected: false,
		},
		{
			String: "%BC",
			Expected: true,
		},
		{
			String: "%Bc",
			Expected: false,
		},
		{
			String: "%BD",
			Expected: true,
		},
		{
			String: "%Bd",
			Expected: false,
		},
		{
			String: "%BE",
			Expected: true,
		},
		{
			String: "%Be",
			Expected: false,
		},
		{
			String: "%BF",
			Expected: true,
		},
		{
			String: "%Bf",
			Expected: false,
		},



		{
			String: "%b0",
			Expected: false,
		},
		{
			String: "%b1",
			Expected: false,
		},
		{
			String: "%b2",
			Expected: false,
		},
		{
			String: "%b3",
			Expected: false,
		},
		{
			String: "%b4",
			Expected: false,
		},
		{
			String: "%b5",
			Expected: false,
		},
		{
			String: "%b6",
			Expected: false,
		},
		{
			String: "%b7",
			Expected: false,
		},
		{
			String: "%b8",
			Expected: false,
		},
		{
			String: "%b9",
			Expected: false,
		},
		{
			String: "%bA",
			Expected: false,
		},
		{
			String: "%ba",
			Expected: false,
		},
		{
			String: "%bB",
			Expected: false,
		},
		{
			String: "%bb",
			Expected: false,
		},
		{
			String: "%bC",
			Expected: false,
		},
		{
			String: "%bc",
			Expected: false,
		},
		{
			String: "%bD",
			Expected: false,
		},
		{
			String: "%bd",
			Expected: false,
		},
		{
			String: "%bE",
			Expected: false,
		},
		{
			String: "%be",
			Expected: false,
		},
		{
			String: "%bF",
			Expected: false,
		},
		{
			String: "%bf",
			Expected: false,
		},



		{
			String: "%C0",
			Expected: true,
		},
		{
			String: "%C1",
			Expected: true,
		},
		{
			String: "%C2",
			Expected: true,
		},
		{
			String: "%C3",
			Expected: true,
		},
		{
			String: "%C4",
			Expected: true,
		},
		{
			String: "%C5",
			Expected: true,
		},
		{
			String: "%C6",
			Expected: true,
		},
		{
			String: "%C7",
			Expected: true,
		},
		{
			String: "%C8",
			Expected: true,
		},
		{
			String: "%C9",
			Expected: true,
		},
		{
			String: "%CA",
			Expected: true,
		},
		{
			String: "%Ca",
			Expected: false,
		},
		{
			String: "%CB",
			Expected: true,
		},
		{
			String: "%Cb",
			Expected: false,
		},
		{
			String: "%CC",
			Expected: true,
		},
		{
			String: "%Cc",
			Expected: false,
		},
		{
			String: "%CD",
			Expected: true,
		},
		{
			String: "%Cd",
			Expected: false,
		},
		{
			String: "%CE",
			Expected: true,
		},
		{
			String: "%Ce",
			Expected: false,
		},
		{
			String: "%CF",
			Expected: true,
		},
		{
			String: "%Cf",
			Expected: false,
		},



		{
			String: "%c0",
			Expected: false,
		},
		{
			String: "%c1",
			Expected: false,
		},
		{
			String: "%c2",
			Expected: false,
		},
		{
			String: "%c3",
			Expected: false,
		},
		{
			String: "%c4",
			Expected: false,
		},
		{
			String: "%c5",
			Expected: false,
		},
		{
			String: "%c6",
			Expected: false,
		},
		{
			String: "%c7",
			Expected: false,
		},
		{
			String: "%c8",
			Expected: false,
		},
		{
			String: "%c9",
			Expected: false,
		},
		{
			String: "%cA",
			Expected: false,
		},
		{
			String: "%ca",
			Expected: false,
		},
		{
			String: "%cB",
			Expected: false,
		},
		{
			String: "%cb",
			Expected: false,
		},
		{
			String: "%cC",
			Expected: false,
		},
		{
			String: "%cc",
			Expected: false,
		},
		{
			String: "%cD",
			Expected: false,
		},
		{
			String: "%cd",
			Expected: false,
		},
		{
			String: "%cE",
			Expected: false,
		},
		{
			String: "%ce",
			Expected: false,
		},
		{
			String: "%cF",
			Expected: false,
		},
		{
			String: "%cf",
			Expected: false,
		},



		{
			String: "%D0",
			Expected: true,
		},
		{
			String: "%D1",
			Expected: true,
		},
		{
			String: "%D2",
			Expected: true,
		},
		{
			String: "%D3",
			Expected: true,
		},
		{
			String: "%D4",
			Expected: true,
		},
		{
			String: "%D5",
			Expected: true,
		},
		{
			String: "%D6",
			Expected: true,
		},
		{
			String: "%D7",
			Expected: true,
		},
		{
			String: "%D8",
			Expected: true,
		},
		{
			String: "%D9",
			Expected: true,
		},
		{
			String: "%DA",
			Expected: true,
		},
		{
			String: "%Da",
			Expected: false,
		},
		{
			String: "%DB",
			Expected: true,
		},
		{
			String: "%Db",
			Expected: false,
		},
		{
			String: "%DC",
			Expected: true,
		},
		{
			String: "%Dc",
			Expected: false,
		},
		{
			String: "%DD",
			Expected: true,
		},
		{
			String: "%Dd",
			Expected: false,
		},
		{
			String: "%DE",
			Expected: true,
		},
		{
			String: "%De",
			Expected: false,
		},
		{
			String: "%DF",
			Expected: true,
		},
		{
			String: "%Df",
			Expected: false,
		},



		{
			String: "%d0",
			Expected: false,
		},
		{
			String: "%d1",
			Expected: false,
		},
		{
			String: "%d2",
			Expected: false,
		},
		{
			String: "%d3",
			Expected: false,
		},
		{
			String: "%d4",
			Expected: false,
		},
		{
			String: "%d5",
			Expected: false,
		},
		{
			String: "%d6",
			Expected: false,
		},
		{
			String: "%d7",
			Expected: false,
		},
		{
			String: "%d8",
			Expected: false,
		},
		{
			String: "%d9",
			Expected: false,
		},
		{
			String: "%dA",
			Expected: false,
		},
		{
			String: "%da",
			Expected: false,
		},
		{
			String: "%dB",
			Expected: false,
		},
		{
			String: "%db",
			Expected: false,
		},
		{
			String: "%dC",
			Expected: false,
		},
		{
			String: "%dc",
			Expected: false,
		},
		{
			String: "%dD",
			Expected: false,
		},
		{
			String: "%dd",
			Expected: false,
		},
		{
			String: "%dE",
			Expected: false,
		},
		{
			String: "%de",
			Expected: false,
		},
		{
			String: "%dF",
			Expected: false,
		},
		{
			String: "%df",
			Expected: false,
		},



		{
			String: "%E0",
			Expected: true,
		},
		{
			String: "%E1",
			Expected: true,
		},
		{
			String: "%E2",
			Expected: true,
		},
		{
			String: "%E3",
			Expected: true,
		},
		{
			String: "%E4",
			Expected: true,
		},
		{
			String: "%E5",
			Expected: true,
		},
		{
			String: "%E6",
			Expected: true,
		},
		{
			String: "%E7",
			Expected: true,
		},
		{
			String: "%E8",
			Expected: true,
		},
		{
			String: "%E9",
			Expected: true,
		},
		{
			String: "%EA",
			Expected: true,
		},
		{
			String: "%Ea",
			Expected: false,
		},
		{
			String: "%EB",
			Expected: true,
		},
		{
			String: "%Eb",
			Expected: false,
		},
		{
			String: "%EC",
			Expected: true,
		},
		{
			String: "%Ec",
			Expected: false,
		},
		{
			String: "%ED",
			Expected: true,
		},
		{
			String: "%Ed",
			Expected: false,
		},
		{
			String: "%EE",
			Expected: true,
		},
		{
			String: "%Ee",
			Expected: false,
		},
		{
			String: "%EF",
			Expected: true,
		},
		{
			String: "%Ef",
			Expected: false,
		},



		{
			String: "%e0",
			Expected: false,
		},
		{
			String: "%e1",
			Expected: false,
		},
		{
			String: "%e2",
			Expected: false,
		},
		{
			String: "%e3",
			Expected: false,
		},
		{
			String: "%e4",
			Expected: false,
		},
		{
			String: "%e5",
			Expected: false,
		},
		{
			String: "%e6",
			Expected: false,
		},
		{
			String: "%e7",
			Expected: false,
		},
		{
			String: "%e8",
			Expected: false,
		},
		{
			String: "%e9",
			Expected: false,
		},
		{
			String: "%eA",
			Expected: false,
		},
		{
			String: "%ea",
			Expected: false,
		},
		{
			String: "%eB",
			Expected: false,
		},
		{
			String: "%eb",
			Expected: false,
		},
		{
			String: "%eC",
			Expected: false,
		},
		{
			String: "%ec",
			Expected: false,
		},
		{
			String: "%eD",
			Expected: false,
		},
		{
			String: "%ed",
			Expected: false,
		},
		{
			String: "%eE",
			Expected: false,
		},
		{
			String: "%ee",
			Expected: false,
		},
		{
			String: "%eF",
			Expected: false,
		},
		{
			String: "%ef",
			Expected: false,
		},



		{
			String: "%F0",
			Expected: true,
		},
		{
			String: "%F1",
			Expected: true,
		},
		{
			String: "%F2",
			Expected: true,
		},
		{
			String: "%F3",
			Expected: true,
		},
		{
			String: "%F4",
			Expected: true,
		},
		{
			String: "%F5",
			Expected: true,
		},
		{
			String: "%F6",
			Expected: true,
		},
		{
			String: "%F7",
			Expected: true,
		},
		{
			String: "%F8",
			Expected: true,
		},
		{
			String: "%F9",
			Expected: true,
		},
		{
			String: "%FA",
			Expected: true,
		},
		{
			String: "%Fa",
			Expected: false,
		},
		{
			String: "%FB",
			Expected: true,
		},
		{
			String: "%Fb",
			Expected: false,
		},
		{
			String: "%FC",
			Expected: true,
		},
		{
			String: "%Fc",
			Expected: false,
		},
		{
			String: "%FD",
			Expected: true,
		},
		{
			String: "%Fd",
			Expected: false,
		},
		{
			String: "%FE",
			Expected: true,
		},
		{
			String: "%Fe",
			Expected: false,
		},
		{
			String: "%FF",
			Expected: true,
		},
		{
			String: "%Ff",
			Expected: false,
		},



		{
			String: "%f0",
			Expected: false,
		},
		{
			String: "%f1",
			Expected: false,
		},
		{
			String: "%f2",
			Expected: false,
		},
		{
			String: "%f3",
			Expected: false,
		},
		{
			String: "%f4",
			Expected: false,
		},
		{
			String: "%f5",
			Expected: false,
		},
		{
			String: "%f6",
			Expected: false,
		},
		{
			String: "%f7",
			Expected: false,
		},
		{
			String: "%f8",
			Expected: false,
		},
		{
			String: "%f9",
			Expected: false,
		},
		{
			String: "%fA",
			Expected: false,
		},
		{
			String: "%fa",
			Expected: false,
		},
		{
			String: "%fB",
			Expected: false,
		},
		{
			String: "%fb",
			Expected: false,
		},
		{
			String: "%fC",
			Expected: false,
		},
		{
			String: "%fc",
			Expected: false,
		},
		{
			String: "%fD",
			Expected: false,
		},
		{
			String: "%fd",
			Expected: false,
		},
		{
			String: "%fE",
			Expected: false,
		},
		{
			String: "%fe",
			Expected: false,
		},
		{
			String: "%fF",
			Expected: false,
		},
		{
			String: "%ff",
			Expected: false,
		},

	}

	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		var hexchars []byte = []byte{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'}
		var chars []byte = []byte{' ','+','-','.','/','0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

		for i:=0; i<0xFF; i++{

			var p []byte

			p = append(p, '%')
			p = append(p, hexchars[randomness.Intn(len(hexchars))])
			p = append(p, hexchars[randomness.Intn(len(hexchars))])

			{
				var limit int = 1 + randomness.Intn(5)

				for j:=0; j<limit; j++ {
					p = append(p, chars[randomness.Intn(len(chars))])
				}
			}


			test := struct{
				String string
				Expected bool
			}{
				String: string(p),
				Expected: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {
		actual := rfc3986.HasPrefixPctEncoded(test.String)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("STRING: %q", test.String)
			continue
		}
	}
}
