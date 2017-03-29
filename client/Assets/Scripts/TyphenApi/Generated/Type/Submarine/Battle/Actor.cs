// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine.Battle
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class Actor : TyphenApi.TypeBase<Actor>, TyphenApi.IRealTimeMessage
    {
        [TyphenApi.QueryStringProperty("id", false)]
        [MessagePack.Key("id")]
        [Newtonsoft.Json.JsonProperty("id")]
        [Newtonsoft.Json.JsonRequired]
        public long Id { get; set; }
        [TyphenApi.QueryStringProperty("user_id", false)]
        [MessagePack.Key("user_id")]
        [Newtonsoft.Json.JsonProperty("user_id")]
        [Newtonsoft.Json.JsonRequired]
        public long UserId { get; set; }
        [TyphenApi.QueryStringProperty("type", false)]
        [MessagePack.Key("type")]
        [Newtonsoft.Json.JsonProperty("type")]
        [Newtonsoft.Json.JsonRequired]
        public TyphenApi.Type.Submarine.Battle.ActorType Type { get; set; }
        [TyphenApi.QueryStringProperty("movement", false)]
        [MessagePack.Key("movement")]
        [Newtonsoft.Json.JsonProperty("movement")]
        [Newtonsoft.Json.JsonRequired]
        public TyphenApi.Type.Submarine.Battle.Movement Movement { get; set; }
        [TyphenApi.QueryStringProperty("is_visible", false)]
        [MessagePack.Key("is_visible")]
        [Newtonsoft.Json.JsonProperty("is_visible")]
        [Newtonsoft.Json.JsonRequired]
        public bool IsVisible { get; set; }
        [TyphenApi.QueryStringProperty("submarine", true)]
        [MessagePack.Key("submarine")]
        [Newtonsoft.Json.JsonProperty("submarine")]
        public TyphenApi.Type.Submarine.Battle.ActorSubmarineObject Submarine { get; set; }
    }
}
