学习笔记

_**作业**_

**我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应
该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？**

_**思路**_
调用标准库或者第三方库的时候，我们需要把error通过wrap处理，保留异常信息，然后返回给上一层，但当我们遇到一个具体的Sentinel errors的时候，
比如 sql.ErrNoRows,我们不应该往上抛出，因为我们依赖的第三方库是有可能会别替换掉的。上层处理Sentinel errors，
当我们的第三方库换成其他第三方库的时候，上层的逻辑就会有问题。所以应该返回的error。就是说，调用的包有Sentinel errors
的时候，Sentinel errors的错误不应该warp。
