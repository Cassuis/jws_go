// automatically generated by the FlatBuffers compiler, do not modify

namespace multiplayMsg
{

using System;
using FlatBuffers;

///玩家状态(数组中以玩家数据中的idx为索引)
public sealed class PlayerState : Table {
  public static PlayerState GetRootAsPlayerState(ByteBuffer _bb) { return GetRootAsPlayerState(_bb, new PlayerState()); }
  public static PlayerState GetRootAsPlayerState(ByteBuffer _bb, PlayerState obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public PlayerState __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  public string AccountID { get { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAccountIDBytes() { return __vector_as_arraysegment(4); }
  ///玩家状态 1-掉线 2-已退出 3-未准备 4-已准备 5-已死亡 6-战斗中
  public int State { get { int o = __offset(6); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  ///玩家hp
  public int GetHp(int j) { int o = __offset(8); return o != 0 ? bb.GetInt(__vector(o) + j * 4) : (int)0; }
  public int HpLength { get { int o = __offset(8); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetHpBytes() { return __vector_as_arraysegment(8); }
  public int Pos { get { int o = __offset(10); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  /// avatar id
  public int GetAvatarID(int j) { int o = __offset(12); return o != 0 ? bb.GetInt(__vector(o) + j * 4) : (int)0; }
  public int AvatarIDLength { get { int o = __offset(12); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetAvatarIDBytes() { return __vector_as_arraysegment(12); }
  /// cur avatar
  public int CurAvatar { get { int o = __offset(14); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }

  public static Offset<PlayerState> CreatePlayerState(FlatBufferBuilder builder,
      StringOffset AccountIDOffset = default(StringOffset),
      int state = 0,
      VectorOffset hpOffset = default(VectorOffset),
      int Pos = 0,
      VectorOffset avatarIDOffset = default(VectorOffset),
      int curAvatar = 0) {
    builder.StartObject(6);
    PlayerState.AddCurAvatar(builder, curAvatar);
    PlayerState.AddAvatarID(builder, avatarIDOffset);
    PlayerState.AddPos(builder, Pos);
    PlayerState.AddHp(builder, hpOffset);
    PlayerState.AddState(builder, state);
    PlayerState.AddAccountID(builder, AccountIDOffset);
    return PlayerState.EndPlayerState(builder);
  }

  public static void StartPlayerState(FlatBufferBuilder builder) { builder.StartObject(6); }
  public static void AddAccountID(FlatBufferBuilder builder, StringOffset AccountIDOffset) { builder.AddOffset(0, AccountIDOffset.Value, 0); }
  public static void AddState(FlatBufferBuilder builder, int state) { builder.AddInt(1, state, 0); }
  public static void AddHp(FlatBufferBuilder builder, VectorOffset hpOffset) { builder.AddOffset(2, hpOffset.Value, 0); }
  public static VectorOffset CreateHpVector(FlatBufferBuilder builder, int[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddInt(data[i]); return builder.EndVector(); }
  public static void StartHpVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static void AddPos(FlatBufferBuilder builder, int Pos) { builder.AddInt(3, Pos, 0); }
  public static void AddAvatarID(FlatBufferBuilder builder, VectorOffset avatarIDOffset) { builder.AddOffset(4, avatarIDOffset.Value, 0); }
  public static VectorOffset CreateAvatarIDVector(FlatBufferBuilder builder, int[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddInt(data[i]); return builder.EndVector(); }
  public static void StartAvatarIDVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static void AddCurAvatar(FlatBufferBuilder builder, int curAvatar) { builder.AddInt(5, curAvatar, 0); }
  public static Offset<PlayerState> EndPlayerState(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<PlayerState>(o);
  }
  public static void FinishPlayerStateBuffer(FlatBufferBuilder builder, Offset<PlayerState> offset) { builder.Finish(offset.Value); }
};


}
