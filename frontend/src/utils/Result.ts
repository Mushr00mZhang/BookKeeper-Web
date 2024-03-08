export interface IResult<T> {
  /**
   * 状态码
   */
  code: number;
  /**
   * 结果
   */
  result: T | null;
  /**
   * 提示
   */
  tip: string;
  /**
   * 错误
   */
  error: any;
}
export type Result<T> =
  | {
      /**
       * 状态码
       */
      code: 0;
      /**
       * 结果
       */
      result: T;
      /**
       * 提示
       */
      tip: string;
      /**
       * 错误
       */
      error: any;
    }
  | {
      /**
       * 状态码
       */
      code: number;
      /**
       * 结果
       */
      result: null;
      /**
       * 提示
       */
      tip: string;
      /**
       * 错误
       */
      error: any;
    };
