# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Battle
  class Vector
    include Virtus.model(:strict => true)

    attribute :x, Number, :required => true
    attribute :y, Number, :required => true
  end
end
