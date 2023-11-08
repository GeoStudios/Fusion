package LinkedList;

// Native Imports
using "#/std";

// return { Index: 0, Next: null, Value: null }
fn Construct() -> HashMap {
    HashMap LinkedList = {
        Index: 0,
        Next: null,
        Value null
    };
    return LinkedList;
}

// returns { Value: WantedValue }
fn GetValue(int index, HashMap list) -> HashMap {
    if (list["Next"] != null && list["Index"] != index) {
        return GetValue(n, list["Next"]);
    }
    if (list["Index"] == index) {
        return { Value: list["Value"] }
    }
    std.print("Err", std.LineSeparator);
    std.exit(0);
}

fn GetLength(HashMap list) -> int {
      (list["Item"] != null) {
        return GetLength(list["Item"]);
    }
    return list["Index"] + 1;
}

// returns string
fn String(HashMap list) -> string {
    int Length = GetLength(list);
    /
    bnz
}