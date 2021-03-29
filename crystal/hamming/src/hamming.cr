# Please implement your solution to hamming in this file

module Hamming
  def self.distance(a, b)
    a.chars.zip(b.chars).count { |x, y| 1 if x != y }
  end
end
