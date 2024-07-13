/// [Leetcode](https://leetcode.cn) solution
pub struct Solution;

use std::collections::HashMap;

impl Solution {
    /// [1. Tow Sum](https://leetcode.cn/problems/two-sum/description)
    ///
    /// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    ///
    /// You may assume that each input would have exactly one solution, and you may not use the same element twice.
    ///
    /// You can return the answer in any order.
    ///
    ///  
    ///
    /// Example 1:
    ///
    //      Input: nums = [2,7,11,15], target = 9
    ///     Output: [0,1]
    ///     Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
    ///
    /// Example 2:
    ///
    ///     Input: nums = [3,2,4], target = 6
    ///     Output: [1,2]
    ///
    /// Example 3:
    ///
    ///     Input: nums = [3,3], target = 6
    ///     Output: [0,1]
    ///     
    ///  
    ///
    /// Constraints:
    ///
    ///     2 <= nums.length <= 104
    ///     -109 <= nums[i] <= 109
    ///     -109 <= target <= 109
    ///     Only one valid answer exists.
    ///
    ///  
    /// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut indices = HashMap::new();
        for (i, v) in nums.iter().enumerate() {
            if let Some(&idx) = indices.get(&(target - v)) {
                return vec![i as i32, idx as i32];
            };
            indices.insert(v, i);
        }
        return vec![];
    }

    /// 49. [Group Anagrams](https://leetcode.cn/problems/group-anagrams/description/)
    ///
    /// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
    ///
    /// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: strs = ["eat","tea","tan","ate","nat","bat"]
    /// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
    ///
    /// Example 2:
    ///
    /// Input: strs = [""]
    /// Output: [[""]]
    ///
    /// Example 3:
    ///
    /// Input: strs = ["a"]
    /// Output: [["a"]]
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     1 <= strs.length <= 104
    ///     0 <= strs[i].length <= 100
    ///     strs[i] consists of lowercase English letters.
    ///
    ///
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut matchers = HashMap::new();
        for str in strs.iter() {
            let mut matcher = [0; 26];
            for e in str.chars() {
                matcher[e as usize - 'a' as usize] += 1;
            }
            let group = matchers.entry(matcher).or_insert(vec![]);
            group.push(str.to_string());
        }

        let mut result: Vec<Vec<String>> = Vec::with_capacity(matchers.len());
        for (_, group) in matchers {
            result.push(group);
        }
        return result;
    }
}
