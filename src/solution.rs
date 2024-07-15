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

    /// [49. Group Anagrams](https://leetcode.cn/problems/group-anagrams/description/)
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

    /// [128. Longest Consecutive Sequence](https://leetcode.cn/problems/longest-consecutive-sequence/description/)
    ///
    /// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
    ///
    /// You must write an algorithm that runs in O(n) time.
    ///
    /// Example 1:
    ///
    /// Input: nums = [100,4,200,1,3,2]
    /// Output: 4
    /// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
    ///
    /// Example 2:
    ///
    /// Input: nums = [0,3,7,2,5,8,4,6,0,1]
    /// Output: 9
    ///
    /// Constraints:
    ///
    ///	0 <= nums.length <= 10^5
    ///	-10^9 <= nums[i] <= 10^9
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut exist: HashMap<i32, bool> = HashMap::new();
        for num in nums {
            exist.insert(num, true);
        }

        let mut longest = 0;
        for (&num, _) in exist.iter() {
            if let Some(&ok) = exist.get(&(num - 1)) {
                if ok {
                    continue;
                }
            }

            let mut length = 1;
            let mut n = num;
            while let Some(&ok) = exist.get(&(n + 1)) {
                if !ok {
                    break;
                }
                length += 1;
                n += 1;
            }
            longest = std::cmp::max(longest, length)
        }
        return longest;
    }

    /// [283. Move Zeroes](https://leetcode.cn/problems/move-zeroes/description/)
    ///
    /// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
    ///
    /// Note that you must do this in-place without making a copy of the array.
    ///
    /// Example 1:
    ///
    /// Input: nums = [0,1,0,3,12]
    /// Output: [1,3,12,0,0]
    ///
    /// Example 2:
    ///
    /// Input: nums = [0]
    /// Output: [0]
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^4
    ///	-2^31 <= nums[i] <= 2^31 - 1
    ///
    /// Follow up: Could you minimize the total number of operations done?
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut tail = 0;
        for head in 0..nums.len() {
            if nums[head] != 0 {
                nums.swap(tail, head);
                tail += 1;
            }
        }
    }

    /// [11. Container with most water](https://leetcode.cn/problems/container-with-most-water/description/)
    ///
    /// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
    ///
    /// Find two lines that together with the x-axis form a container, such that the container contains the most water.
    ///
    /// Return the maximum amount of water a container can store.
    ///
    /// Notice that you may not slant the container.
    ///
    ///  
    ///
    /// Example 1:
    ///
    ///  8|                    |                |              
    ///  7|                    |                |              |
    ///  6|          |         |          |     |              |
    ///  5|          |         |          |     |       |      |
    ///  4|          |         |          |     |       |      |
    ///  3|          |         |          |     |       |      |
    ///  2|          |         |    |     |     |       |      |
    ///  1|    |     |         |    |     |     |       |      |
    ///   -------------------------------------------------------
    ///   0    1     2         3    4     5     6       7      8
    ///
    /// Input: height = [1,8,6,2,5,4,8,3,7]
    /// Output: 49
    /// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
    ///
    /// Example 2:
    ///
    /// Input: height = [1,1]
    /// Output: 1
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     n == height.length
    ///     2 <= n <= 105
    ///     0 <= height[i] <= 104
    ///
    ///
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut i = 0;
        let mut j = height.len() - 1;
        let mut max_area = 0;

        while j > i {
            if height[i] > height[j] {
                let area = height[j] * (j - i) as i32;
                max_area = max_area.max(area);
                j -= 1;
            } else {
                let area = height[i] * (j - i) as i32;
                max_area = max_area.max(area);
                i += 1;
            }
        }
        return max_area;
    }
}
