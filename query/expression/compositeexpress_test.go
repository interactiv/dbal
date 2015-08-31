package expression_test

import (
	"testing"

	"github.com/interactiv/expect"
	"github.com/mparaiso/dbal/query/expression"
)

func TestCompositeExpression(t *testing.T) {
	e := expect.New(t)
	expr := &expression.CompositeExpression{expression.OR, []string{"u.group_id = 1"}}
	e.Expect(expr.Length()).ToBe(1)
	expr.Add("u.group_id = 2")
	e.Expect(expr.Length()).ToBe(2)
}

//<?php

//namespace Doctrine\Tests\DBAL\Query\Expression;

//use Doctrine\DBAL\Query\Expression\CompositeExpression;

//    /**
//     * @dataProvider provideDataForConvertToString
//     */
//    public function testCompositeUsageAndGeneration($type, $parts, $expects)
//    {
//        $expr = new CompositeExpression($type, $parts);

//        $this->assertEquals($expects, (string) $expr);
//    }

//    public function provideDataForConvertToString()
//    {
//        return array(
//            array(
//                CompositeExpression::TYPE_AND,
//                array('u.user = 1'),
//                'u.user = 1'
//            ),
//            array(
//                CompositeExpression::TYPE_AND,
//                array('u.user = 1', 'u.group_id = 1'),
//                '(u.user = 1) AND (u.group_id = 1)'
//            ),
//            array(
//                CompositeExpression::TYPE_OR,
//                array('u.user = 1'),
//                'u.user = 1'
//            ),
//            array(
//                CompositeExpression::TYPE_OR,
//                array('u.group_id = 1', 'u.group_id = 2'),
//                '(u.group_id = 1) OR (u.group_id = 2)'
//            ),
//            array(
//                CompositeExpression::TYPE_AND,
//                array(
//                    'u.user = 1',
//                    new CompositeExpression(
//                        CompositeExpression::TYPE_OR,
//                        array('u.group_id = 1', 'u.group_id = 2')
//                    )
//                ),
//                '(u.user = 1) AND ((u.group_id = 1) OR (u.group_id = 2))'
//            ),
//            array(
//                CompositeExpression::TYPE_OR,
//                array(
//                    'u.group_id = 1',
//                    new CompositeExpression(
//                        CompositeExpression::TYPE_AND,
//                        array('u.user = 1', 'u.group_id = 2')
//                    )
//                ),
//                '(u.group_id = 1) OR ((u.user = 1) AND (u.group_id = 2))'
//            ),
//        );
//    }
//}
