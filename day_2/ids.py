class IDChecker:
  # def __init__(self) -> None:
    # self.ids: list[str] = []

  def check_id(self, id: str) -> [str]:
    invalid_ids: list[str] = []
    first_id, second_id = id.split('-')
    print(f"Checking ids: {int(first_id)} to {int(second_id)}")

    for i in range(int(first_id), int(second_id) + 1):
      # print(f"Checking id: {i}")
      if (len(str(i)) % 2 != 0):
        # Odd number of digits - not possible to repeat
        next
      else:
        string_i = str(i)
        left_digits = string_i[0:len(string_i)//2]
        right_digits = string_i[len(string_i)//2:]
        if left_digits == right_digits:
          invalid_ids.append(int(string_i))

    return invalid_ids

def get_input(filename: str) -> list[str]:
    input: str = ''
    input_list: list[str] = []
    with open(filename, "r") as input_file:
      input = input_file.read()
      input_list = input.split(',')
    return input_list

# input = get_input('./test_input.txt')
input = get_input('./input.txt')
print(input)


running_total: int = 0
for id in input:
  result = IDChecker().check_id(id)
  print(result)
  running_total += sum(result)

print(f"Running total: {running_total}")

