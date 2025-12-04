class Batteries:
  def largest(self, options: str, exclude_index: int ) -> [int, int]:
    options_list = list(options)

    possible_answers = []

    if exclude_index is not None:
      possible_answers = options_list[exclude_index + 1:]
    else:
      possible_answers = options_list[:-1]
    print(f"Possible answers: {possible_answers}")

    return max(possible_answers), possible_answers.index(max(possible_answers))

def get_input(filename: str) -> list[str]:
    input_list: list[str] = []
    with open(filename, "r") as input_file:
      input = input_file.read()
      input_list = input.split('\n')
    return input_list

input = [line for line in get_input('./input.txt') if line]
print(input)

running_total = 0

for index, line in enumerate(input):
  largest, largest_index = Batteries().largest(line, None)
  print(f"Largest: {largest} at index {largest_index}")
  second_largest, second_largest_index = Batteries().largest(line, largest_index)
  print(f"Second largest: {second_largest}")
  print (f"Answer:{largest}{second_largest}")

  running_total += int(f"{largest}{second_largest}")
  print('---')

print(f"Running total: {running_total}")
