DIGITS = 12

class Batteries:
  def largest(self, aggregator: [str], remaining_options: [str] ) -> [[str], [str]]:

    if (len(aggregator) == DIGITS) or (len(remaining_options) == 0):
      return [aggregator, remaining_options]
    else:
      int_options = [int(option) for option in remaining_options]

      if (len(aggregator) - DIGITS) + 1 == 0:
        possible_options = int_options
      else:
        possible_options = int_options[0: ((len(aggregator) - DIGITS) + 1)]

      max_option = max(possible_options)
      max_option_index = remaining_options.index(str(max_option))
      return self.largest(aggregator + [str(max_option)], remaining_options[max_option_index + 1:])

def get_input(filename: str) -> list[str]:
    input_list: list[str] = []
    with open(filename, "r") as input_file:
      input = input_file.read()
      input_list = input.split('\n')
    return input_list

input = [line for line in get_input('./input.txt') if line]
print(input)
print('---')

running_total = 0
for line in input:
  # print(f"Line: {line}")
  answer = ''.join(Batteries().largest([], line)[0])
  print(f"Answer: {answer}")
  running_total += int(answer)
  print('---')

print(f"Running total: {running_total}")


888911112111
888911112111
888191111211

3121910778619
3121910778619
