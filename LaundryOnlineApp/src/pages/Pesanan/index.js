import React from 'react'
import { StyleSheet, Text, View } from 'react-native'
import { Searchbar } from 'react-native-paper';
import { WARNA_ABU_ABU, WARNA_UTAMA } from '../../utils/constant'

const Pesanan = () => {
    const [searchQuery, setSearchQuery] = React.useState('');

    const onChangeSearch = query => setSearchQuery(query);

    return (
        <View style={styles.page}>
            <View style={styles.searchBox}>
                <Searchbar
                    style={styles.searchBar}
                    placeholder="Masukkan no. pesanan"
                    onChangeText={onChangeSearch}
                    value={searchQuery}
                    inputStyle={styles.placeholder}
                />
            </View>
        </View>
    )
}

export default Pesanan

const styles = StyleSheet.create({
    page: {
        flex: 1,
        backgroundColor: WARNA_ABU_ABU,
    },
    searchBox: {
        backgroundColor: WARNA_UTAMA,
        paddingHorizontal: 20,
        paddingVertical: 10
    },
    searchBar:{
        borderRadius: 50,
        height: 40,
    },
    placeholder: {
        fontSize: 18,
        fontFamily: 'Roboto',
    }
})
