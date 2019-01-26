# Food Planner
A simple command line script to help generate our menus each week.

This utility takes in a set of JSON files:
 * ```inventory.json``` - A list of locations of food stuffs, and the count for each item.
 * ```nutrition.json``` - Key-Value lookup of the nutritional information for each item.
 * ```plan.json``` - The template plan of what meals should be generated, and roughly how large (e.g. meal, snack, treat) and for how many individuals