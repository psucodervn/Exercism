# Please implement your solution to rna-transcription in this file

module RnaComplement
  def self.of_dna(s)
    m = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }
    s.chars.map { |x| m[x] }.join
  end
end
