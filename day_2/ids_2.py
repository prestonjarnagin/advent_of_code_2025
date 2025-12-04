class IDChecker:
  # def __init__(self) -> None:
    # self.ids: list[str] = []

  def check_id_range(self, id: str) -> [str]:
    invalid_ids: list[str] = []
    first_id, second_id = id.split('-')
    print(f"Checking ids: {int(first_id)} to {int(second_id)}")

    for id in range(int(first_id), int(second_id) + 1):
      print(f">> Checking id: {id}")

      # possible_sequence_lengths = self.find_possible_sequence_lengths(id)
      # print(possible_sequence_lengths)
      # for possibility in possible_sequence_lengths:
        # id_string = str(id) # '123123123'

        # pattern = id_string[0:possibility] # '123' for 3
        # segments = [id_string[x:x+possibility] for x in range(0, len(id_string), possibility)]
        # => ['123', '123', '123']

        # self.segments_equal(segments)
    # return invalid_ids

      max_sequence_length = len(str(id))//2
      possible_sequence_lengths = list(range(1, max_sequence_length + 1))
      print(f"Possible sequence lengths: {possible_sequence_lengths}")

      for sequence_length in possible_sequence_lengths:
        segments = self.create_segments(str(id), sequence_length)
        print(f"Segments: {segments}")
        if self.segments_equal(segments):
          print(f"Segment equal ✅: {segments}")
          invalid_ids.append(id)
          break
        else:
          print(f"Segment not equal ❌: {segments}")

    return invalid_ids

  @staticmethod
  def create_segments(id: str, sequence_length: int) -> [str]:
    return [id[x:x+sequence_length] for x in range(0, len(id), sequence_length)]

  @staticmethod
  def segments_equal(segments: [str]) -> bool:
    for subsegment in segments:
      if subsegment != segments[0]:
        return False

    # print(f"Segments are equal: {segments}")
    return True

  # @staticmethod
  # def find_possible_sequence_lengths(id: int) -> [int]:
  #   ceiling = id//2
  #   possible_range = range(1, ceiling + 1)
  #   possible_sequence_lengths = []

  #   for i in possible_range:
  #     # print(i)
  #     if id % i == 0:
  #       possible_sequence_lengths.append(i)
  #   return possible_sequence_lengths


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
for id_range in input:
  result = IDChecker().check_id_range(id_range)
  print(result)
  running_total += sum(result)

print(f"Running total: {running_total}")

# print(IDChecker.find_possible_sequence_lengths(1))
# print(IDChecker.segments_equal(['12', '12']))
# print(IDChecker.create_segments('123123123', 3))

