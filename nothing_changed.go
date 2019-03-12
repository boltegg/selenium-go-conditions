package conditions

import (
	"github.com/tebeka/selenium"
	"time"
)

func NothingChanged() selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {

		dom, err := getDom(wd)
		if err != nil {
			return false, err
		}

		time.Sleep(time.Second * 2)

		dom2, err := getDom(wd)
		if err != nil {
			return false, err
		}

		if dom != dom2 {
			return false, nil
		}

		return true, nil
	}
}

//

var jsdom = `var getDOM = (function() {
    var dom = "",
        depth = 0;
    
    return function(node, n) {
        for (var i = 0; i < depth; i++) {
            dom += '<span>|---</span>';
        }
        
        dom += '<b>' + node.nodeName.toLowerCase() + '</b>';

        if (node.id) {
            dom += '[#' + node.id + ']';
        }

        if (node.className) {
            dom += '(' + node.className + ')'
        }
               
        if (typeof n === 'number') {
            dom += '<span>{child #' + n + '}</span>';
        }

        dom += '<br>';
        depth++;

        [].forEach.call(node.children, function(node, childNumber) {
            getDOM(node, childNumber);
        });

        depth--;
        return dom;
    };
})();

return getDOM(document.body);
`

func getDom(wd selenium.WebDriver) (string, error) {
	res, err := wd.ExecuteScript(jsdom, nil)
	if err != nil {
		return "", err
	}

	s, _ := res.(string)
	return s, nil
}