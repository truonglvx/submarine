// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class GetRoomsObject : TyphenApi.TypeBase<GetRoomsObject>
    {
        protected static readonly SerializationInfo<GetRoomsObject, List<Room>> rooms = new SerializationInfo<GetRoomsObject, List<Room>>("rooms", false, (x) => x.Rooms, (x, v) => x.Rooms = v);
        public List<Room> Rooms { get; set; }
    }
}
