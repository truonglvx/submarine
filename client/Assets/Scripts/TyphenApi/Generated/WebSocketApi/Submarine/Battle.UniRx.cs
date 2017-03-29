// This file was generated by typhen-api

using System;
using UniRx;

namespace TyphenApi.WebSocketApi.Parts.Submarine
{
    public partial class Battle : TyphenApi.IWebSocketApi
    {
        public IObservable<TyphenApi.Type.Submarine.Battle.Ping> OnPingReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Ping>(h => OnPingReceive += h, h => OnPingReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Room> OnRoomReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Room>(h => OnRoomReceive += h, h => OnRoomReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Now> OnNowReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Now>(h => OnNowReceive += h, h => OnNowReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Start> OnStartReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Start>(h => OnStartReceive += h, h => OnStartReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Finish> OnFinishReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Finish>(h => OnFinishReceive += h, h => OnFinishReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Actor> OnActorReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Actor>(h => OnActorReceive += h, h => OnActorReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Visibility> OnVisibilityReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Visibility>(h => OnVisibilityReceive += h, h => OnVisibilityReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Movement> OnMovementReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Movement>(h => OnMovementReceive += h, h => OnMovementReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Destruction> OnDestructionReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Destruction>(h => OnDestructionReceive += h, h => OnDestructionReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Pinger> OnPingerReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Pinger>(h => OnPingerReceive += h, h => OnPingerReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.Equipment> OnEquipmentReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.Equipment>(h => OnEquipmentReceive += h, h => OnEquipmentReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.StartRequest> OnStartRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.StartRequest>(h => OnStartRequestReceive += h, h => OnStartRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.AccelerationRequest> OnAccelerationRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.AccelerationRequest>(h => OnAccelerationRequestReceive += h, h => OnAccelerationRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.BrakeRequest> OnBrakeRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.BrakeRequest>(h => OnBrakeRequestReceive += h, h => OnBrakeRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.TurnRequest> OnTurnRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.TurnRequest>(h => OnTurnRequestReceive += h, h => OnTurnRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.PingerRequest> OnPingerRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.PingerRequest>(h => OnPingerRequestReceive += h, h => OnPingerRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.TorpedoRequest> OnTorpedoRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.TorpedoRequest>(h => OnTorpedoRequestReceive += h, h => OnTorpedoRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.WatcherRequest> OnWatcherRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.WatcherRequest>(h => OnWatcherRequestReceive += h, h => OnWatcherRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.AddBotRequest> OnAddBotRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.AddBotRequest>(h => OnAddBotRequestReceive += h, h => OnAddBotRequestReceive -= h);
        }
        public IObservable<TyphenApi.Type.Submarine.Battle.RemoveBotRequest> OnRemoveBotRequestReceiveAsObservable()
        {
            return Observable.FromEvent<TyphenApi.Type.Submarine.Battle.RemoveBotRequest>(h => OnRemoveBotRequestReceive += h, h => OnRemoveBotRequestReceive -= h);
        }
    }
}
