# Please implement your solution to perfect-numbers in this file

module PerfectNumbers
  def self.classify(n)
    raise ArgumentError.new if n <= 0
    d = (1..n - 1).select { |x| n % x == 0 }.sum
    if d == n
      "perfect"
    elsif d > n
      "abundant"
    else
      "deficient"
    end
  end
end
