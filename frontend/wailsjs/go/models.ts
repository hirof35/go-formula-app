export namespace main {
	
	export class EvaluateResult {
	    result: string;
	    isErr: boolean;
	
	    static createFrom(source: any = {}) {
	        return new EvaluateResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.result = source["result"];
	        this.isErr = source["isErr"];
	    }
	}

}

