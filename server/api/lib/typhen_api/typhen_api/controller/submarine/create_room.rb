# This file was generated by typhen-api

module TyphenApi::Controller::Submarine
  module CreateRoom
    extend ActiveSupport::Concern

    class RequestType
      include Virtus.model(:strict => true)

    end

    ResponseType = TyphenApi::Model::Submarine::CreateRoomObject
    ErrorType = TyphenApi::Model::Submarine::Error

    def no_authentication_required?
      false
    end
  end
end
