// Code generated by scripts/generate_go_funcs; DO NOT EDIT.
package jsoncfg

import "strings"

func (c *JSONCfg) getNestedKey(key string) interface{} {
	var keys = strings.Split(key, ".")
	var val interface{} = c.config

	for _, key := range keys {
		val = val.(map[string]interface{})[key]
	}

	return val
}

func (c *JSONCfg) Get(key string) interface{} {
	return c.getNestedKey(key)
}

func (c *JSONCfg) GetArray(key string) []interface{} {
	return c.getNestedKey(key).([]interface{})
}

func (c *JSONCfg) GetMap(key string) map[string]interface{} {
	return c.getNestedKey(key).(map[string]interface{})
}

func (c *JSONCfg) GetDiff(key string) interface{} {
	return c.getNestedKey(key)
}

func (c *JSONCfg) GetDiffArray(key string) []interface{} {
	return c.getNestedKey(key).([]interface{})
}

func (c *JSONCfg) GetDiffMap(key string) map[string]interface{} {
	return c.getNestedKey(key).(map[string]interface{})
}

func (c *JSONCfg) GetBool(key string) bool {
	return c.getNestedKey(key).(bool)
}

func (c *JSONCfg) GetBoolArray(key string) []bool {
	return c.getNestedKey(key).([]bool)
}

func (c *JSONCfg) GetBoolMap(key string) map[string]bool {
	return c.getNestedKey(key).(map[string]bool)
}

func (c *JSONCfg) GetDiffBool(key string) bool {
	return c.getNestedKey(key).(bool)
}

func (c *JSONCfg) GetDiffBoolArray(key string) []bool {
	return c.getNestedKey(key).([]bool)
}

func (c *JSONCfg) GetDiffBoolMap(key string) map[string]bool {
	return c.getNestedKey(key).(map[string]bool)
}

func (c *JSONCfg) GetFloat32(key string) float32 {
	return float32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetFloat32Array(key string) []float32 {
	var ok bool
	var val = []float32{}

	if val, ok = c.getNestedKey(key).([]float32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, float32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetFloat32Map(key string) map[string]float32 {
	var ok bool
	var val = map[string]float32{}

	if val, ok = c.getNestedKey(key).(map[string]float32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = float32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffFloat32(key string) float32 {
	return float32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffFloat32Array(key string) []float32 {
	var ok bool
	var val = []float32{}

	if val, ok = c.getNestedKey(key).([]float32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, float32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffFloat32Map(key string) map[string]float32 {
	var ok bool
	var val = map[string]float32{}

	if val, ok = c.getNestedKey(key).(map[string]float32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = float32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetFloat64(key string) float64 {
	return float64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetFloat64Array(key string) []float64 {
	var ok bool
	var val = []float64{}

	if val, ok = c.getNestedKey(key).([]float64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, float64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetFloat64Map(key string) map[string]float64 {
	var ok bool
	var val = map[string]float64{}

	if val, ok = c.getNestedKey(key).(map[string]float64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = float64(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffFloat64(key string) float64 {
	return float64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffFloat64Array(key string) []float64 {
	var ok bool
	var val = []float64{}

	if val, ok = c.getNestedKey(key).([]float64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, float64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffFloat64Map(key string) map[string]float64 {
	var ok bool
	var val = map[string]float64{}

	if val, ok = c.getNestedKey(key).(map[string]float64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = float64(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetInt(key string) int {
	return int(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetIntArray(key string) []int {
	var ok bool
	var val = []int{}

	if val, ok = c.getNestedKey(key).([]int); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetIntMap(key string) map[string]int {
	var ok bool
	var val = map[string]int{}

	if val, ok = c.getNestedKey(key).(map[string]int); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffInt(key string) int {
	return int(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffIntArray(key string) []int {
	var ok bool
	var val = []int{}

	if val, ok = c.getNestedKey(key).([]int); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffIntMap(key string) map[string]int {
	var ok bool
	var val = map[string]int{}

	if val, ok = c.getNestedKey(key).(map[string]int); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetInt16(key string) int16 {
	return int16(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetInt16Array(key string) []int16 {
	var ok bool
	var val = []int16{}

	if val, ok = c.getNestedKey(key).([]int16); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int16(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetInt16Map(key string) map[string]int16 {
	var ok bool
	var val = map[string]int16{}

	if val, ok = c.getNestedKey(key).(map[string]int16); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int16(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffInt16(key string) int16 {
	return int16(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffInt16Array(key string) []int16 {
	var ok bool
	var val = []int16{}

	if val, ok = c.getNestedKey(key).([]int16); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int16(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffInt16Map(key string) map[string]int16 {
	var ok bool
	var val = map[string]int16{}

	if val, ok = c.getNestedKey(key).(map[string]int16); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int16(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetInt32(key string) int32 {
	return int32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetInt32Array(key string) []int32 {
	var ok bool
	var val = []int32{}

	if val, ok = c.getNestedKey(key).([]int32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetInt32Map(key string) map[string]int32 {
	var ok bool
	var val = map[string]int32{}

	if val, ok = c.getNestedKey(key).(map[string]int32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffInt32(key string) int32 {
	return int32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffInt32Array(key string) []int32 {
	var ok bool
	var val = []int32{}

	if val, ok = c.getNestedKey(key).([]int32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffInt32Map(key string) map[string]int32 {
	var ok bool
	var val = map[string]int32{}

	if val, ok = c.getNestedKey(key).(map[string]int32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetInt64(key string) int64 {
	return int64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetInt64Array(key string) []int64 {
	var ok bool
	var val = []int64{}

	if val, ok = c.getNestedKey(key).([]int64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetInt64Map(key string) map[string]int64 {
	var ok bool
	var val = map[string]int64{}

	if val, ok = c.getNestedKey(key).(map[string]int64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int64(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffInt64(key string) int64 {
	return int64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffInt64Array(key string) []int64 {
	var ok bool
	var val = []int64{}

	if val, ok = c.getNestedKey(key).([]int64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, int64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffInt64Map(key string) map[string]int64 {
	var ok bool
	var val = map[string]int64{}

	if val, ok = c.getNestedKey(key).(map[string]int64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = int64(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetString(key string) string {
	return c.getNestedKey(key).(string)
}

func (c *JSONCfg) GetStringArray(key string) []string {
	var ok bool
	var val = []string{}

	if val, ok = c.getNestedKey(key).([]string); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, v.(string))
	}

	return val
}

func (c *JSONCfg) GetStringMap(key string) map[string]string {
	var ok bool
	var val = map[string]string{}

	if val, ok = c.getNestedKey(key).(map[string]string); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = v.(string)
	}

	return val
}

func (c *JSONCfg) GetDiffString(key string) string {
	return c.getNestedKey(key).(string)
}

func (c *JSONCfg) GetDiffStringArray(key string) []string {
	var ok bool
	var val = []string{}

	if val, ok = c.getNestedKey(key).([]string); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, v.(string))
	}

	return val
}

func (c *JSONCfg) GetDiffStringMap(key string) map[string]string {
	var ok bool
	var val = map[string]string{}

	if val, ok = c.getNestedKey(key).(map[string]string); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = v.(string)
	}

	return val
}

func (c *JSONCfg) GetUint(key string) uint {
	return uint(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetUintArray(key string) []uint {
	var ok bool
	var val = []uint{}

	if val, ok = c.getNestedKey(key).([]uint); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetUintMap(key string) map[string]uint {
	var ok bool
	var val = map[string]uint{}

	if val, ok = c.getNestedKey(key).(map[string]uint); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffUint(key string) uint {
	return uint(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffUintArray(key string) []uint {
	var ok bool
	var val = []uint{}

	if val, ok = c.getNestedKey(key).([]uint); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffUintMap(key string) map[string]uint {
	var ok bool
	var val = map[string]uint{}

	if val, ok = c.getNestedKey(key).(map[string]uint); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetUint16(key string) uint16 {
	return uint16(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetUint16Array(key string) []uint16 {
	var ok bool
	var val = []uint16{}

	if val, ok = c.getNestedKey(key).([]uint16); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint16(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetUint16Map(key string) map[string]uint16 {
	var ok bool
	var val = map[string]uint16{}

	if val, ok = c.getNestedKey(key).(map[string]uint16); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint16(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffUint16(key string) uint16 {
	return uint16(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffUint16Array(key string) []uint16 {
	var ok bool
	var val = []uint16{}

	if val, ok = c.getNestedKey(key).([]uint16); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint16(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffUint16Map(key string) map[string]uint16 {
	var ok bool
	var val = map[string]uint16{}

	if val, ok = c.getNestedKey(key).(map[string]uint16); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint16(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetUint32(key string) uint32 {
	return uint32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetUint32Array(key string) []uint32 {
	var ok bool
	var val = []uint32{}

	if val, ok = c.getNestedKey(key).([]uint32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetUint32Map(key string) map[string]uint32 {
	var ok bool
	var val = map[string]uint32{}

	if val, ok = c.getNestedKey(key).(map[string]uint32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffUint32(key string) uint32 {
	return uint32(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffUint32Array(key string) []uint32 {
	var ok bool
	var val = []uint32{}

	if val, ok = c.getNestedKey(key).([]uint32); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint32(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffUint32Map(key string) map[string]uint32 {
	var ok bool
	var val = map[string]uint32{}

	if val, ok = c.getNestedKey(key).(map[string]uint32); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint32(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetUint64(key string) uint64 {
	return uint64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetUint64Array(key string) []uint64 {
	var ok bool
	var val = []uint64{}

	if val, ok = c.getNestedKey(key).([]uint64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetUint64Map(key string) map[string]uint64 {
	var ok bool
	var val = map[string]uint64{}

	if val, ok = c.getNestedKey(key).(map[string]uint64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint64(v.(float64))
	}

	return val
}

func (c *JSONCfg) GetDiffUint64(key string) uint64 {
	return uint64(c.getNestedKey(key).(float64))
}

func (c *JSONCfg) GetDiffUint64Array(key string) []uint64 {
	var ok bool
	var val = []uint64{}

	if val, ok = c.getNestedKey(key).([]uint64); ok {
		return val
	}

	for _, v := range c.getNestedKey(key).([]interface{}) {
		val = append(val, uint64(v.(float64)))
	}

	return val
}

func (c *JSONCfg) GetDiffUint64Map(key string) map[string]uint64 {
	var ok bool
	var val = map[string]uint64{}

	if val, ok = c.getNestedKey(key).(map[string]uint64); ok {
		return val
	}

	for k, v := range c.getNestedKey(key).(map[string]interface{}) {
		val[k] = uint64(v.(float64))
	}

	return val
}
