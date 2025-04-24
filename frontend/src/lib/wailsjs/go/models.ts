export namespace functions {
	
	export class CPUStats {
	    usage: number;
	    timestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new CPUStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.usage = source["usage"];
	        this.timestamp = source["timestamp"];
	    }
	}
	export class PermissionStatus {
	    IsElevated: boolean;
	    RequiresElevation: boolean;
	    UnaccessiblePaths: string[];
	    ElevationCommand: string;
	    CanCleanUserFiles: boolean;
	    CanCleanSystemFiles: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PermissionStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IsElevated = source["IsElevated"];
	        this.RequiresElevation = source["RequiresElevation"];
	        this.UnaccessiblePaths = source["UnaccessiblePaths"];
	        this.ElevationCommand = source["ElevationCommand"];
	        this.CanCleanUserFiles = source["CanCleanUserFiles"];
	        this.CanCleanSystemFiles = source["CanCleanSystemFiles"];
	    }
	}
	export class CleanerResult {
	    Files: {[key: string]: FileInfo[]};
	    TotalSize: number;
	    Permissions: PermissionStatus;
	
	    static createFrom(source: any = {}) {
	        return new CleanerResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Files = source["Files"];
	        this.TotalSize = source["TotalSize"];
	        this.Permissions = this.convertValues(source["Permissions"], PermissionStatus);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class FileInfo {
	    Path: string;
	    Size: number;
	    Name: string;
	    Location: string;
	    NeedsElevation: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Size = source["Size"];
	        this.Name = source["Name"];
	        this.Location = source["Location"];
	        this.NeedsElevation = source["NeedsElevation"];
	    }
	}
	export class MemoryStats {
	    totalRAMMB: number;
	    freeRAMMB: number;
	    usedRAMMB: number;
	    usagePercent: number;
	    timestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new MemoryStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totalRAMMB = source["totalRAMMB"];
	        this.freeRAMMB = source["freeRAMMB"];
	        this.usedRAMMB = source["usedRAMMB"];
	        this.usagePercent = source["usagePercent"];
	        this.timestamp = source["timestamp"];
	    }
	}

}

