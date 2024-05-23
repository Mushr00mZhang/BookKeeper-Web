/**
 * 获取分页列表Dto
 */
export interface IGetDto {
  /**
   * 序号
   */
  index: number;
  /**
   * 大小
   */
  size: number;
}
/**
 * 分页列表Dto
 */
export interface IDto<T> {
  /**
   * 总数
   */
  total: number;
  /**
   * 列表
   */
  items: T[];
}
