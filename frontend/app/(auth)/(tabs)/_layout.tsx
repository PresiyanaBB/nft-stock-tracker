import { TabBarIcon } from '@/components/navigation/TabBarIcon';
import { useAuth } from '@/context/AuthContext';
import { UserRole } from '@/types/user';
import { Ionicons } from '@expo/vector-icons';
import { Tabs } from 'expo-router';
import { ComponentProps } from 'react';
import { Text } from 'react-native';

export default function TabLayout() {
  const { user } = useAuth();

  const tabs = [
    {
      showFor: [UserRole.Collector, UserRole.Admin],
      name: '(buy-nft)',
      displayName: 'Buy NFT',
      icon: 'search',
      options: {
        headerShown: false
      }
    },
    {
      showFor: [UserRole.Collector],
      name: '(owned-nft)',
      displayName: 'Owned NFT',
      icon: 'wallet',
      options: {
        headerShown: false
      }
    },
    {
      showFor: [UserRole.Admin],
      name: 'scan-nft',
      displayName: 'Scan NFT',
      icon: 'qr-code',
      options: {
        headerShown: true
      }
    },
    {
      showFor: [UserRole.Collector, UserRole.Admin],
      name: '(stock)',
      displayName: 'Stocks',
      icon: 'bar-chart',
      options: {
        headerShown: false,
      }
    },
    {
      showFor: [UserRole.Collector, UserRole.Admin],
      name: 'settings',
      displayName: 'Settings',
      icon: 'cog',
      options: {
        headerShown: true,
      }
    }
  ];

  return (
    <Tabs>
      {tabs.map(tab => (
        <Tabs.Screen
          key={tab.name}
          name={tab.name}
          options={{
            ...tab.options,
            headerTitle: tab.displayName,
            href: tab.showFor.includes(user?.role!) ? tab.name : null,
            tabBarLabel: ({ focused }) => (
              <Text style={{ color: focused ? "black" : "gray", fontSize: 12 }} >
                {tab.displayName}
              </Text>
            ),
            tabBarIcon: ({ focused }) => (
              <TabBarIcon
                name={tab.icon as ComponentProps<typeof Ionicons>['name']}
                color={focused ? 'black' : "gray"}
              />
            )
          }}
        />
      ))}
    </Tabs>
  );
}