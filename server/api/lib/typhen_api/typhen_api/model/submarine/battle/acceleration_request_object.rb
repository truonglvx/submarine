# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Battle
  class AccelerationRequestObject
    include Virtus.model(:strict => true)

    attribute :direction, Number, :required => true
  end
end
