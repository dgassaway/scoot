/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SuppressWarnings({"cast", "rawtypes", "serial", "unchecked"})
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2018-01-10")
public class Command implements org.apache.thrift.TBase<Command, Command._Fields>, java.io.Serializable, Cloneable, Comparable<Command> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("Command");

  private static final org.apache.thrift.protocol.TField ARGV_FIELD_DESC = new org.apache.thrift.protocol.TField("argv", org.apache.thrift.protocol.TType.LIST, (short)1);
  private static final org.apache.thrift.protocol.TField ENV_VARS_FIELD_DESC = new org.apache.thrift.protocol.TField("envVars", org.apache.thrift.protocol.TType.MAP, (short)2);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new CommandStandardSchemeFactory());
    schemes.put(TupleScheme.class, new CommandTupleSchemeFactory());
  }

  public List<String> argv; // required
  public Map<String,String> envVars; // optional

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    ARGV((short)1, "argv"),
    ENV_VARS((short)2, "envVars");

    private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

    static {
      for (_Fields field : EnumSet.allOf(_Fields.class)) {
        byName.put(field.getFieldName(), field);
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, or null if its not found.
     */
    public static _Fields findByThriftId(int fieldId) {
      switch(fieldId) {
        case 1: // ARGV
          return ARGV;
        case 2: // ENV_VARS
          return ENV_VARS;
        default:
          return null;
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, throwing an exception
     * if it is not found.
     */
    public static _Fields findByThriftIdOrThrow(int fieldId) {
      _Fields fields = findByThriftId(fieldId);
      if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
      return fields;
    }

    /**
     * Find the _Fields constant that matches name, or null if its not found.
     */
    public static _Fields findByName(String name) {
      return byName.get(name);
    }

    private final short _thriftId;
    private final String _fieldName;

    _Fields(short thriftId, String fieldName) {
      _thriftId = thriftId;
      _fieldName = fieldName;
    }

    public short getThriftFieldId() {
      return _thriftId;
    }

    public String getFieldName() {
      return _fieldName;
    }
  }

  // isset id assignments
  private static final _Fields optionals[] = {_Fields.ENV_VARS};
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.ARGV, new org.apache.thrift.meta_data.FieldMetaData("argv", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.ListMetaData(org.apache.thrift.protocol.TType.LIST, 
            new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING))));
    tmpMap.put(_Fields.ENV_VARS, new org.apache.thrift.meta_data.FieldMetaData("envVars", org.apache.thrift.TFieldRequirementType.OPTIONAL, 
        new org.apache.thrift.meta_data.MapMetaData(org.apache.thrift.protocol.TType.MAP, 
            new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING), 
            new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING))));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(Command.class, metaDataMap);
  }

  public Command() {
  }

  public Command(
    List<String> argv)
  {
    this();
    this.argv = argv;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public Command(Command other) {
    if (other.isSetArgv()) {
      List<String> __this__argv = new ArrayList<String>(other.argv);
      this.argv = __this__argv;
    }
    if (other.isSetEnvVars()) {
      Map<String,String> __this__envVars = new HashMap<String,String>(other.envVars);
      this.envVars = __this__envVars;
    }
  }

  public Command deepCopy() {
    return new Command(this);
  }

  @Override
  public void clear() {
    this.argv = null;
    this.envVars = null;
  }

  public int getArgvSize() {
    return (this.argv == null) ? 0 : this.argv.size();
  }

  public java.util.Iterator<String> getArgvIterator() {
    return (this.argv == null) ? null : this.argv.iterator();
  }

  public void addToArgv(String elem) {
    if (this.argv == null) {
      this.argv = new ArrayList<String>();
    }
    this.argv.add(elem);
  }

  public List<String> getArgv() {
    return this.argv;
  }

  public Command setArgv(List<String> argv) {
    this.argv = argv;
    return this;
  }

  public void unsetArgv() {
    this.argv = null;
  }

  /** Returns true if field argv is set (has been assigned a value) and false otherwise */
  public boolean isSetArgv() {
    return this.argv != null;
  }

  public void setArgvIsSet(boolean value) {
    if (!value) {
      this.argv = null;
    }
  }

  public int getEnvVarsSize() {
    return (this.envVars == null) ? 0 : this.envVars.size();
  }

  public void putToEnvVars(String key, String val) {
    if (this.envVars == null) {
      this.envVars = new HashMap<String,String>();
    }
    this.envVars.put(key, val);
  }

  public Map<String,String> getEnvVars() {
    return this.envVars;
  }

  public Command setEnvVars(Map<String,String> envVars) {
    this.envVars = envVars;
    return this;
  }

  public void unsetEnvVars() {
    this.envVars = null;
  }

  /** Returns true if field envVars is set (has been assigned a value) and false otherwise */
  public boolean isSetEnvVars() {
    return this.envVars != null;
  }

  public void setEnvVarsIsSet(boolean value) {
    if (!value) {
      this.envVars = null;
    }
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case ARGV:
      if (value == null) {
        unsetArgv();
      } else {
        setArgv((List<String>)value);
      }
      break;

    case ENV_VARS:
      if (value == null) {
        unsetEnvVars();
      } else {
        setEnvVars((Map<String,String>)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case ARGV:
      return getArgv();

    case ENV_VARS:
      return getEnvVars();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case ARGV:
      return isSetArgv();
    case ENV_VARS:
      return isSetEnvVars();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof Command)
      return this.equals((Command)that);
    return false;
  }

  public boolean equals(Command that) {
    if (that == null)
      return false;

    boolean this_present_argv = true && this.isSetArgv();
    boolean that_present_argv = true && that.isSetArgv();
    if (this_present_argv || that_present_argv) {
      if (!(this_present_argv && that_present_argv))
        return false;
      if (!this.argv.equals(that.argv))
        return false;
    }

    boolean this_present_envVars = true && this.isSetEnvVars();
    boolean that_present_envVars = true && that.isSetEnvVars();
    if (this_present_envVars || that_present_envVars) {
      if (!(this_present_envVars && that_present_envVars))
        return false;
      if (!this.envVars.equals(that.envVars))
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_argv = true && (isSetArgv());
    list.add(present_argv);
    if (present_argv)
      list.add(argv);

    boolean present_envVars = true && (isSetEnvVars());
    list.add(present_envVars);
    if (present_envVars)
      list.add(envVars);

    return list.hashCode();
  }

  @Override
  public int compareTo(Command other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetArgv()).compareTo(other.isSetArgv());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetArgv()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.argv, other.argv);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetEnvVars()).compareTo(other.isSetEnvVars());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetEnvVars()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.envVars, other.envVars);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    return 0;
  }

  public _Fields fieldForId(int fieldId) {
    return _Fields.findByThriftId(fieldId);
  }

  public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
    schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
  }

  public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
    schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder("Command(");
    boolean first = true;

    sb.append("argv:");
    if (this.argv == null) {
      sb.append("null");
    } else {
      sb.append(this.argv);
    }
    first = false;
    if (isSetEnvVars()) {
      if (!first) sb.append(", ");
      sb.append("envVars:");
      if (this.envVars == null) {
        sb.append("null");
      } else {
        sb.append(this.envVars);
      }
      first = false;
    }
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws org.apache.thrift.TException {
    // check for required fields
    // check for sub-struct validity
  }

  private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
    try {
      write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
    try {
      read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private static class CommandStandardSchemeFactory implements SchemeFactory {
    public CommandStandardScheme getScheme() {
      return new CommandStandardScheme();
    }
  }

  private static class CommandStandardScheme extends StandardScheme<Command> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, Command struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // ARGV
            if (schemeField.type == org.apache.thrift.protocol.TType.LIST) {
              {
                org.apache.thrift.protocol.TList _list0 = iprot.readListBegin();
                struct.argv = new ArrayList<String>(_list0.size);
                String _elem1;
                for (int _i2 = 0; _i2 < _list0.size; ++_i2)
                {
                  _elem1 = iprot.readString();
                  struct.argv.add(_elem1);
                }
                iprot.readListEnd();
              }
              struct.setArgvIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // ENV_VARS
            if (schemeField.type == org.apache.thrift.protocol.TType.MAP) {
              {
                org.apache.thrift.protocol.TMap _map3 = iprot.readMapBegin();
                struct.envVars = new HashMap<String,String>(2*_map3.size);
                String _key4;
                String _val5;
                for (int _i6 = 0; _i6 < _map3.size; ++_i6)
                {
                  _key4 = iprot.readString();
                  _val5 = iprot.readString();
                  struct.envVars.put(_key4, _val5);
                }
                iprot.readMapEnd();
              }
              struct.setEnvVarsIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          default:
            org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
        }
        iprot.readFieldEnd();
      }
      iprot.readStructEnd();

      // check for required fields of primitive type, which can't be checked in the validate method
      struct.validate();
    }

    public void write(org.apache.thrift.protocol.TProtocol oprot, Command struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      if (struct.argv != null) {
        oprot.writeFieldBegin(ARGV_FIELD_DESC);
        {
          oprot.writeListBegin(new org.apache.thrift.protocol.TList(org.apache.thrift.protocol.TType.STRING, struct.argv.size()));
          for (String _iter7 : struct.argv)
          {
            oprot.writeString(_iter7);
          }
          oprot.writeListEnd();
        }
        oprot.writeFieldEnd();
      }
      if (struct.envVars != null) {
        if (struct.isSetEnvVars()) {
          oprot.writeFieldBegin(ENV_VARS_FIELD_DESC);
          {
            oprot.writeMapBegin(new org.apache.thrift.protocol.TMap(org.apache.thrift.protocol.TType.STRING, org.apache.thrift.protocol.TType.STRING, struct.envVars.size()));
            for (Map.Entry<String, String> _iter8 : struct.envVars.entrySet())
            {
              oprot.writeString(_iter8.getKey());
              oprot.writeString(_iter8.getValue());
            }
            oprot.writeMapEnd();
          }
          oprot.writeFieldEnd();
        }
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class CommandTupleSchemeFactory implements SchemeFactory {
    public CommandTupleScheme getScheme() {
      return new CommandTupleScheme();
    }
  }

  private static class CommandTupleScheme extends TupleScheme<Command> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, Command struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      BitSet optionals = new BitSet();
      if (struct.isSetArgv()) {
        optionals.set(0);
      }
      if (struct.isSetEnvVars()) {
        optionals.set(1);
      }
      oprot.writeBitSet(optionals, 2);
      if (struct.isSetArgv()) {
        {
          oprot.writeI32(struct.argv.size());
          for (String _iter9 : struct.argv)
          {
            oprot.writeString(_iter9);
          }
        }
      }
      if (struct.isSetEnvVars()) {
        {
          oprot.writeI32(struct.envVars.size());
          for (Map.Entry<String, String> _iter10 : struct.envVars.entrySet())
          {
            oprot.writeString(_iter10.getKey());
            oprot.writeString(_iter10.getValue());
          }
        }
      }
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, Command struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      BitSet incoming = iprot.readBitSet(2);
      if (incoming.get(0)) {
        {
          org.apache.thrift.protocol.TList _list11 = new org.apache.thrift.protocol.TList(org.apache.thrift.protocol.TType.STRING, iprot.readI32());
          struct.argv = new ArrayList<String>(_list11.size);
          String _elem12;
          for (int _i13 = 0; _i13 < _list11.size; ++_i13)
          {
            _elem12 = iprot.readString();
            struct.argv.add(_elem12);
          }
        }
        struct.setArgvIsSet(true);
      }
      if (incoming.get(1)) {
        {
          org.apache.thrift.protocol.TMap _map14 = new org.apache.thrift.protocol.TMap(org.apache.thrift.protocol.TType.STRING, org.apache.thrift.protocol.TType.STRING, iprot.readI32());
          struct.envVars = new HashMap<String,String>(2*_map14.size);
          String _key15;
          String _val16;
          for (int _i17 = 0; _i17 < _map14.size; ++_i17)
          {
            _key15 = iprot.readString();
            _val16 = iprot.readString();
            struct.envVars.put(_key15, _val16);
          }
        }
        struct.setEnvVarsIsSet(true);
      }
    }
  }

}

