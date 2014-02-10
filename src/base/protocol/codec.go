package protocol

import ()

/**
 * 编码解码接口
 * @author abram
 * @since 0.0.1
 */
type ICodec interface {
	/**
	 * 解码
	 * @param buf socket 接受到的二进制数据
	 * @author abram
	 * @return protoPack 数据包对象
	 * @return err 错误
	 */
	Decode(buf []byte) (protoPack *ProtoPack, err error)

	/**
	 * 编码
	 * @param protoPack 数据包对象
	 * @author abram
	 * @return buf 二进制数据流
	 * @return err
	 */
	Encode(protoPack *ProtoPack) (buf []byte, err error)
}

/**
 * 默认的编码解码器
 * @author abram
 */
type DefaultCodec struct {
	isReadHead bool
	length     int
	cacheData  []byte
}

func NewDefaultCodec() *DefaultCodec {
	return &DefaultCodec{}
}

func (codec *DefaultCodec) Decode(buf []byte) (protoPack *ProtoPack, err error) {
	if !codec.isReadHead {
		if len(buf) < 4 {

		}
	}
	return nil, nil
}

func (codec *DefaultCodec) Encode(protoPack ProtoPack) (buf []byte, err error) {
	return nil, nil
}
