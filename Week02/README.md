学习笔记

_**作业**_

**我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应
该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？**

_**思路**_
调用标准库或者第三方库的时候，我们需要把errors通过wrap处理，保留异常信息，然后抛给上一层。
但是当我们遇到sql.ErrNoRows的时候，不应该直接wrap往上抛。具体做法应该是warp自定义sentinel errprs往上抛。
1.有对应的响应api错因为查询为空，误，与500不服务内部错误不同，就应该直接nil加warp自定义sentinel errprs错误往上抛，最后is自定义的sentinel errprs。
2.上层调用不应该先判errors是否为空，再判断查询记录为空。
3.如果wrap后，在api层做is或者as的时候，可能会使用了sql.ErrNoRows，来返回调用errors。