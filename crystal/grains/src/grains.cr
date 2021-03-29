# Please implement your solution to grains in this file

module Grains
  def self.square(x)
    raise ArgumentError.new if x > 64 || x <= 0
    2_u64 ** (x - 1)
  end

  def self.total
    2_u64 ** 63 - 1 + 2_u64 ** 63
  end
end
