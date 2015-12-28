# This file was generated by typhen-api

Rails.application.routes.draw do

  post 'ping' => 'ping#service'
  post 'sign_up' => 'sign_up#service'
  post 'login' => 'login#service'
  post 'find_user' => 'find_user#service'
  post 'create_room' => 'create_room#service'
  post 'get_rooms' => 'get_rooms#service'
  post 'join_into_room' => 'join_into_room#service'
  post 'battle/find_room' => 'battle/find_room#service'
  post 'battle/close_room' => 'battle/close_room#service'

end
