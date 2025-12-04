class Cypher
  def initialize
    @current_place = 50
    @max_place = 99

    @zero_counter = 0
    @zero_pass_counter = 0
  end

  def process(instructions)
    instructions.each do |instruction|

      puts '---'

      direction = instruction[0]
      value = instruction[1]

      puts "Step: #{direction}#{value}"
      step(direction: direction, value: value)
      puts @current_place

      next unless instruction[2]

      if instruction[2] == @current_place
        puts "✅"
      else
        puts "❌"
      end

      next unless instruction[3]

      if instruction[3] == @zero_pass_counter
        puts "Zero: ✅"
      else
        puts "Zero: ❌"
      end
    end

      puts '--'
      puts '--'
      puts '--'

    puts "Zero Counter: #{@zero_counter}"
    puts "Zero Pass Counter: #{@zero_pass_counter}"

    puts "Total: #{@zero_counter + @zero_pass_counter}"
  end

  def step(direction:, value:)
    case direction
    when 'L'
      calculate_l_passes(value)
      calculate_l(value)
    when 'R'
      calculate_r_passes(value)
      calculate_r(value)
    end

    @zero_counter += 1 if @current_place == 0
  end

  private

  def calculate_l_passes(value)
    return if @current_place.zero? && value < (@max_place + 2)

    return unless (@current_place - value).negative?

    # Don't count landing on zero as a 'pass'
    # return if ().abs == 100

    number_of_passes = (value - @current_place).abs / (@max_place + 1) + 1

    @zero_pass_counter += number_of_passes
    puts "Updated zero pass counter: #{@zero_pass_counter}"
  end

  def calculate_r_passes(value)
    return unless (value + @current_place) > (@max_place + 1)

    # Don't count landing on zero as a 'pass'
    # puts "Value: #{value}"
    # puts "Current place: #{@current_place}"
    # puts "Max place: #{@max_place}"
    # puts "Value + current place: #{value + @current_place}"
    return if value + @current_place == 100

    number_of_passes = (value + @current_place) / (@max_place + 1)
    # puts "Number of passes: #{number_of_passes}"

    @zero_pass_counter += number_of_passes
    puts "Updated zero pass counter: #{@zero_pass_counter}"
  end

  def calculate_l(value)
    @current_place = ((@current_place % (@max_place + 1)) - value) % (@max_place + 1)
  end

  def calculate_r(value)
    @current_place = ((@current_place % (@max_place + 1)) + value) % (@max_place + 1)
  end
end

cc = Cypher.new
cc.process([
             ['L', 68, 82, 1],
             ['L', 30, 52, 1],
             ['R', 48, 0, 1],
             ['L', 5, 95, 1],
             ['R', 60, 55, 2],
             ['L', 55, 0, 2],
             ['L', 1, 99, 2],
             ['L', 99, 0, 2],
             ['R', 14, 14, 2],
             ['L', 82, 32, 3],
            #  ['L', 1000, 32, 13],
            #  ['R', 1000, 32, 23]
           ])

cc = Cypher.new
arr = [
             ['L', 50, 0, 0],
             ['R', 50, 50, 0],
             ['R', 50, 0, 0],
             ['L', 100, 0, 0],
             ['R', 101, 1, 1]

           ]

cc.process(arr)

input_path = File.join(File.dirname(__FILE__), "input.txt")
lines = File.readlines(input_path).map(&:strip).reject(&:empty?)
instructions = lines.map { |line| [line[0], line[1..].to_i] }

cc = Cypher.new
results = cc.process(instructions)
