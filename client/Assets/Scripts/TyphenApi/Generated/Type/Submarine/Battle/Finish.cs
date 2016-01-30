// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Finish : TyphenApi.TypeBase<Finish>
    {
        protected static readonly SerializationInfo<Finish, bool> hasWon = new SerializationInfo<Finish, bool>("has_won", false, (x) => x.HasWon, (x, v) => x.HasWon = v);
        public bool HasWon { get; set; }
        protected static readonly SerializationInfo<Finish, long> finishedAt = new SerializationInfo<Finish, long>("finished_at", false, (x) => x.FinishedAt, (x, v) => x.FinishedAt = v);
        public long FinishedAt { get; set; }
    }
}
