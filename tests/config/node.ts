import {AllParameter, MyPackage,TestEnum} from './out/js'
import * as skip from './out/js_skip'
import {MyMessage} from './out/js_skip/mypackage'
import * as DepDep from './out/js_skip/depdep'

let p = new AllParameter()
p.valueBool = true;
p.valueInt32 = 25;
p.valueTestEnum = TestEnum.SNAKE_TEST_VALUE;

let buf = p.pack()
let d = new AllParameter(buf);
if (d.valueBool != p.valueBool || d.valueInt32 != p.valueInt32 || d.valueTestEnum != p.valueTestEnum) {
	console.log(p)
	console.log(d)
	throw Error("serialize error")
}

new skip.DependMessage();
new MyMessage();
new DepDep.DependTestMessage();

console.log("TypeScript Success")

