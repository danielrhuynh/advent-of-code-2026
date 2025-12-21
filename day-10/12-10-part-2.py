import re
from dataclasses import dataclass
from typing import List, Tuple

from z3 import Int, Solver, Sum, sat


@dataclass
class Machine:
    diagram: str
    buttons: List[List[int]]
    req: List[int]

# Porting over my parsing from go
def parse_input(path: str) -> List[Machine]:
    machines: List[Machine] = []
    with open(path, "r", encoding="utf-8") as f:
        for raw in f:
            line = raw.strip()
            if not line:
                continue
            lb = line.find('[')
            rb = line.find(']')

            diagram = line[lb+1:rb]

            rest = line[rb+1:]
            cb = rest.find('{')

            buttons_part = rest[:cb]

            buttons: List[List[int]] = []
            s = buttons_part
            while True:
                ob = s.find('(')
                if ob == -1:
                    break
                close = s.find(')', ob+1)

                inner = s[ob+1:close].strip()
                if inner == "":
                    btn = []
                else:
                    parts = [p.strip() for p in inner.split(',')]
                    btn = [int(p) for p in parts if p != ""]
                buttons.append(btn)

                s = s[close+1:]
            lcur = line.find('{')
            rcur = line.rfind('}')

            req_str = line[lcur+1:rcur].strip()
            if req_str == "":
                req = []
            else:
                req_parts = [p.strip() for p in req_str.split(',')]
                req = [int(p) for p in req_parts if p != ""]
            machines.append(Machine(diagram=diagram, buttons=buttons, req=req))
    return machines


def solve_joltage(machine: Machine) -> Tuple[int, List[int]]:
    n = len(machine.req)
    m = len(machine.buttons)

    # affects[i] = list of buttons that increment counter i
    # this is almost like an adjacency list where we map counters to the buttons that affect them
    affects = [[] for _ in range(n)]
    for j, btn in enumerate(machine.buttons):
        for i in btn:
            affects[i].append(j)

    # Defining unknown variables
    x = [Int(f"x_{j}") for j in range(m)]

    s = Solver()

    # adds constraints:
    # x cannot be negative
    # each variable can only be as high as their upper bound (to limit search space)
    for j, btn in enumerate(machine.buttons):
        s.add(x[j] >= 0)
        ub = min(machine.req[i] for i in btn)
        s.add(x[j] <= ub)

    # adds the actual constraint of hitting the count exactly
    for i in range(n):
        s.add(Sum([x[j] for j in affects[i]]) == machine.req[i])

    # the objective is the total number of button presses
    total = Sum(x)

    # check if s is feasible
    if s.check() != sat:
        raise ValueError("UNSAT: no way to meet requirements exactly with given buttons")

    # use binary search to find the minimum possible number of presses
    # basically what is binary search does is it checks for a total num of presses m by temporarily adding a constraint
    # if a feasible solution exists, if a solution exists, we know that we can't make the solution
    # infeasible by adding more presses so we bring r = m
    # otherwise if the solution is infeasible, we have too little presses so we set l = m+1
    # we do this until we converge on a solution
    hi = sum(machine.req)
    lo = 0

    best_model = None
    best_total = None

    while lo < hi:
        mid = (lo + hi) // 2
        s.push()
        s.add(total <= mid)
        ok = (s.check() == sat)
        if ok:
            hi = mid
            best_model = s.model()
            best_total = mid
        else:
            lo = mid + 1
        s.pop()

    # add total <= low as a permanent constraint and generate a feasible model and extract the button counts which we sum and return
    s.add(total <= lo)
    assert s.check() == sat
    model = s.model()

    sol = [model.evaluate(xj, model_completion=True).as_long() for xj in x]
    min_presses = sum(sol)
    return min_presses, sol

def main(path: str) -> None:
    machines = parse_input(path)

    grand_total = 0
    for k, machine in enumerate(machines):
        presses, sol = solve_joltage(machine)
        grand_total += presses

    print("answer:", grand_total)


if __name__ == "__main__":
    main("day-10/input.txt")
